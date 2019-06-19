<template>
    <div :class="'absolute right-0 top-0' + (open ? ' w-1/4' : '')">
        <div @click="Toggle()" class="cursor-pointer bg-black text-white p-4" v-html="(open) ? '>' : '<'"></div>
        <div :class="'w-full h-full bg-black text-white p-4 pr-10 ' + (open ? '' : 'hidden')">
            <div>
                <div v-for="(product, index) in processedProducts" :key="index" class="flex flex-1 w-full">
                    <div class="w-3/5 cursor-pointer" @click="OpenProduct(product)">{{ product.Name }} x{{ product.Quantity ? product.Quantity : 1 }}</div><div class="w-1/4">&euro; {{product.Price.toFixed(2)}}</div><div class="cursor-pointer" @click="IncreaseQuantity(product.UUID)">+</div>&nbsp;<div class="cursor-pointer" @click="RemoveProduct(product.UUID)">-</div>
                </div>
            </div>
            <div class="border-0 border-t border-solid border-white flex flex-1 w-full mt-4 pt-2">
                <div class="w-3/4">Total:</div><div class="w-1/4">&euro; {{store.state.total_value.toFixed(2)}}</div>
            </div>
        </div>
    </div>
</template>

<script>
import Store from '@/../store/index.js'

export default {
    name: 'Cart',

    data() {
        return {
            open: false,
            products: Store.state.products,
            processedProducts: [],
            store: Store
        }
    },

    mounted() {
        this.ProcessProducts()
    },

    methods: {
        OpenProduct(ev) {
            this.$router.push('/product/' + ev.srcElement.id)
        },

        RemoveProduct(UUID) {
            Store.commit('Remove', UUID)
            this.value = Store.state.total_value
            this.products = Store.state.products
        },

        Toggle() {
            this.open = !this.open
        },

        IncreaseQuantity(UUID) {
            for(let i = 0; i < Store.state.products.length; i++) {
                if(Store.state.products[i].UUID == UUID) {
                    Store.commit('Push', Object.assign({}, Store.state.products[i]))
                    return
                }
            }
        },

        ProcessProducts() {
            this.processedProducts = []
            for(let i = 0; i < this.products.length; i++) {
                let exists = false
                for(let j = 0; j < this.processedProducts.length; j++) {
                    if(this.products[i].UUID == this.processedProducts[j].UUID) {
                        exists = true
                        if(this.processedProducts[j].Quantity == undefined) {
                            this.processedProducts[j].Quantity = 1
                        }
                        this.processedProducts[j].Quantity += 1
                        break
                    }
                }
                if(!exists) {
                    this.processedProducts.push(Object.assign({},this.products[i]))
                }
            }
        }
    },

    watch: {
        products() {
            this.ProcessProducts()
        }
    }
}
</script>

