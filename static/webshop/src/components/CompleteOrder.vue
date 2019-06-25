<template>
  <div class="w-full h-full">
    <cart ref="cart"/>
    <navbar/>
    <div>
        <div class="flex flex-1 w-2/3 mx-auto mt-12 border border-solid border-gray-300 p-4">
            <div class="w-1/2">
                <div class="flex flex-1">
                    Thank you! Your order has been registered and will be worked on as soon as possible. We will keep you updated through your email.
                </div>
            </div>
        </div>
    </div>
  </div>
</template>

<script>
import Navbar from '@/components/objects/Navbar'
import Cart from '@/components/objects/Cart'
import Store from '@/../store/index.js'

export default {
  name: 'AccountPanel',

  data() {
    return {
      user_info: {}
    }
  },

  components: {
    'navbar': Navbar,
    'cart': Cart
  },

  mounted() {
    if(Store.state.products.length > 0) {
      axios.post('/api/order', {
        products: JSON.stringify(Store.state.products)
      }).then(response => {
        if(response.data.success == true) {
          Store.commit('Clear')
          Cart.products = []
        }
      })
    }
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

