<template>
  <div class="w-full h-full">
    <navbar/>
    <div class="flex flex-1 flex-wrap">
        <div v-for="(product, index) in products" :key="index" class="product my-2 h-32 border-b border-solid border-0 border-gray-200">
            {{ product.Name }}
            {{ GetImage(product.Images[0]) }}
            <span class="text-gray-400">{{ product.Description }}</span>
        </div>
    </div>
  </div>
</template>

<script>
import Navbar from '@/components/objects/Navbar'

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
        GetImage(imageId) {
            //axios.get('')
        }
    },

    components: {
        'navbar': Navbar
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


