<template>
    <div class="absolute right-0 h-screen w-1/6 bg-black text-white p-4">
        <div>
            <div v-for="(product, index) in products" :key="index" class="flex flex-1 w-full">
                <div class="w-3/4 cursor-pointer" @click="OpenProduct(product)">{{product.Name}}</div><div class="w-1/4">{{product.Price}}</div><div class="cursor-pointer" @click="RemoveProduct(index)">-</div>
            </div>
        </div>
        <div class="border-0 border-t border-solid border-white flex flex-1 w-full mt-4 pt-2">
            <div class="w-3/4">Total:</div><div class="w-1/4">{{store.state.total_value}}</div>
        </div>
    </div>
</template>

<script>
import Store from '@/../store/index.js'

export default {
    name: 'Cart',

    data() {
        return {
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
        }
    }
}
</script>

