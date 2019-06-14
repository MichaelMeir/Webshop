<template>
    <div :class="'absolute right-0 top-0' + (open ? ' w-1/4' : '')">
        <div @click="Toggle()" class="cursor-pointer bg-black text-white p-4" v-html="(open) ? '>' : '<'"></div>
        <div :class="'w-full h-full bg-black text-white p-4 pr-10 ' + (open ? '' : 'hidden')">
            <div>
                <div v-for="(product, index) in products" :key="index" class="flex flex-1 w-full">
                    <div class="w-3/5 cursor-pointer" @click="OpenProduct(product)">{{product.Name}}</div><div class="w-1/4">&euro; {{product.Price.toFixed(2)}}</div><div class="cursor-pointer" @click="RemoveProduct(index)">-</div>
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
            store: Store
        }
    },

    methods: {
        OpenProduct(ev) {
            this.$router.push('/product/' + ev.srcElement.id)
        },

        RemoveProduct(index) {
            Store.commit('Remove', index)
            this.value = Store.state.total_value
            this.products = Store.state.products
        },

        Toggle() {
            this.open = !this.open
        }
    }
}
</script>

