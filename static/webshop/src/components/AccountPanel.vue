<template>
  <div class="w-full h-full">
    <cart/>
    <navbar/>
    <div>
        <div class="flex flex-1 w-2/3 mx-auto mt-12 border border-solid border-gray-300 p-4">
            <div class="w-1/2">
                <h2>Account info</h2>
                <div class="flex flex-1">
                    <div class="w-1/2">
                        <p class="m-0 p-3">username:</p>
                        <p class="m-0 p-3">email:</p>
                        <p class="m-0 p-3">address:</p>
                        <p class="m-0 p-3">zip code:</p>
                        <p class="m-0 p-3">province:</p>
                        <p class="m-0 p-3">country:</p>
                    </div>
                    <div class="w-1/2">
                        <input disabled class="p-2 h-4 mt-2 w-4/5 border-0 bg-white outline-none border-b border-solid border-black" type="text" v-model="user_info.Name"><br>
                        <input @change="updateInfo" class="p-2 h-4 mt-2 w-4/5 border-0 bg-white outline-none border-b border-solid border-black" type="text" v-model="user_info.Email"><br>
                        <input @change="updateInfo" class="p-2 h-4 mt-2 w-4/5 border-0 bg-white outline-none border-b border-solid border-black" type="text" v-model="user_info.Address"><br>
                        <input @change="updateInfo" class="p-2 h-4 mt-2 w-4/5 border-0 bg-white outline-none border-b border-solid border-black" type="text" v-model="user_info.Zip_code"><br>
                        <input @change="updateInfo" class="p-2 h-4 mt-2 w-4/5 border-0 bg-white outline-none border-b border-solid border-black" type="text" v-model="user_info.Province"><br>
                        <input @change="updateInfo" class="p-2 h-4 mt-2 w-4/5 border-0 bg-white outline-none border-b border-solid border-black" type="text" v-model="user_info.Country"><br>
                    </div>
                </div>
            </div>
        </div>
    </div>
  </div>
</template>

<script>
import Navbar from '@/components/objects/Navbar'
import Cart from '@/components/objects/Cart'

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
      axios.get('/api/auth/info').then(response => {
          this.user_info = response.data
      })
  },

  methods: {
      updateInfo() {
          axios.post('/api/auth/update', this.user_info).then(response => {
              console.log(response.data)
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

