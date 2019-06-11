<template>
  <div class="w-full h-full">
    <cart/>
    <navbar/>
    <div class="flex flex-1 w-full h-full">
        <div class="w-1/2 h-full overflow-hidden flex flex-1">
          <product-image class="w-full m-auto" ref="previewImage" :load="loaded" :src="this.product.Images[0]"/>
          <ul class="flex flex-1 p-0 absolute bottom-0">
            <li v-for="(image, index) in this.product.Images" :key="index" class="flex flex-1 w-20 h-20 overflow-hidden rounded-full border p-0 border-black border-solid" @click="setPreviewImage">
              <product-image class="w-full cursor-pointer my-auto" :load="loaded" :src="image"/>
            </li>
          </ul>
        </div>
        <div class="w-1/2 p-4">
          <div class="flex flex-1">
            <h2 class="p-2 bg-black text-white">{{product.Name}}</h2>
          </div>
          <div>{{product.Description}}</div>
          <div class="flex flex-1">
            <div class="cursor-pointer bg-black text-white p-2 mt-5" @click="AddProduct">Order for &euro; {{ product.Price.toFixed(2) }}</div>
          </div>
        </div>
    </div>
  </div>
</template>

<script>
import Navbar from '@/components/objects/Navbar'
import Image from '@/components/objects/ProductImage'
import Cart from '@/components/objects/Cart'

import Store from '@/../store/index.js'

export default {
  name: 'Index',

  data() {
    return {
      product: {
        Price: 0,
        Images: []
      },
      loaded: false
    }
  },

  components: {
    'navbar': Navbar,
    'product-image': Image,
    'cart': Cart
  },

  beforeCreate() {
    axios.get('http://localhost/api/products/specific?id=' + this.$route.params.uuid).then(response => {
      response.data.Images = JSON.parse(response.data.Images)
      this.product = response.data
      console.log(this.product)
      this.loaded = true
    })
  },

  methods: {
    setPreviewImage(ev) {
      this.$refs.previewImage.SetImage(ev.srcElement.id)
    },

    AddProduct() {
      Store.commit('Push', this.product)
    }
  }
}
</script>
