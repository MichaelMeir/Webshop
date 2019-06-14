<template>
  <div class="w-full h-full">
    <cart/>
    <navbar/>
    
  </div>
</template>

<script>
import Navbar from '@/components/objects/Navbar'
import Cart from '@/components/objects/Cart'

export default {
  name: 'AccountPanel',

  data() {
    return {
      categories: []
    }
  },

  components: {
    'navbar': Navbar,
    'cart': Cart
  },
}
</script>

<style>
.category {
    width: 21%;
    margin-left: 1%;
    margin-right: 1%;
}
</style>

