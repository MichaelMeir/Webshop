<template>
  <div class="w-full h-full">
      <cart/>
    <navbar/>
    <div class="flex flex-1 flex-wrap">
        <div v-for="(product, index) in products" :key="index" class="product my-2 h-64 border rounded border-solid border-0 hover:border-gray-400 border-gray-200">
            <product-tile :product="product" :image_id="JSON.parse(product.Images)[0]"></product-tile>
        </div>
    </div>
  </div>
</template>

<script>
import Navbar from '@/components/objects/Navbar'
import ProductTile from '@/components/objects/ProductTile'
import Cart from '@/components/objects/Cart'

export default {
    name: 'CategoryPage',

    mounted() {
        axios.get('/api/categories/specific?name=' + this.$route.params.name).then(response => {
            this.products = response.data.Products
        })
    },

    data() {
        return {
            products: []
        }
    },

    methods: {
    },

    components: {
        'navbar': Navbar,
        'product-tile': ProductTile,
        'cart': Cart
    }
}
</script>

<style>
.product {
    width: 21%;
    margin-left: 1%;
    margin-right: 1%;
}
</style>


