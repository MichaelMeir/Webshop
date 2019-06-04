<template>
  <div class="w-full h-full">
    <navbar/>
    <div>
        {{product.Name}}
        <product-image :load="loaded" :src="this.product.Images[0]"/>
    </div>
  </div>
</template>

<script>
import Navbar from '@/components/objects/Navbar'
import Image from '@/components/objects/ProductImage'

export default {
  name: 'Index',

  data() {
    return {
      product: {
          Images: []
      },
      loaded: false
    }
  },

  components: {
    'navbar': Navbar,
    'product-image': Image
  },

  beforeCreate() {
    axios.get('http://localhost/api/products/specific?id=' + this.$route.params.uuid).then(response => {
      response.data.Images = JSON.parse(response.data.Images)
      this.product = response.data
      this.loaded = true
    })
  }
}
</script>
