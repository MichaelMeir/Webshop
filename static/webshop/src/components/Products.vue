<template>
  <div class="w-full h-full">
    <navbar/>
    <div class="flex flex-1 flex-wrap">
        <div v-for="(product, index) in products" :key="index" class="product my-2 h-32 border-b border-solid border-0 border-gray-200">
            <product-tile :product="product" :image_id="JSON.parse(product.Images)[0]"></product-tile>
        </div>
    </div>
  </div>
</template>

<script>
import Navbar from '@/components/objects/Navbar'
import ProductTile from '@/components/objects/ProductTile'

export default {
    name: 'Index',

    mounted() {
        axios.get('http://localhost/api/products/all').then(response => {
            this.products = response.data
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
        'product-tile': ProductTile
    }
}
</script>

<style>
.product {
    width: 21%;
    margin-left: 2%;
    margin-right: 2%;
}
</style>


