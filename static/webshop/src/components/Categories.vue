<template>
  <div class="w-full h-full">
    <cart/>
    <navbar/>
    <div v-for="(category, index) in categories" :key="index" class="category mt-2">
      <category :category="category"></category>
    </div>
  </div>
</template>

<script>
import Navbar from '@/components/objects/Navbar'
import CategoryTile from '@/components/objects/CategoryTile'
import Cart from '@/components/objects/Cart'

export default {
  name: 'Index',

  data() {
    return {
      categories: []
    }
  },

  components: {
    'navbar': Navbar,
    'category': CategoryTile,
    'cart': Cart
  },

  mounted() {
    axios.get('/api/categories/all').then(response => {
      this.categories = response.data
    })
  }
}
</script>

<style>
.category {
    width: 21%;
    margin-left: 1%;
    margin-right: 1%;
}
</style>

