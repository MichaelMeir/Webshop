<template>
    <div>
        <cart/>
        <navbar/>
        <div>
            <div class="bg-black text-white" v-if="err.length != 0">{{err}}</div>
            <div class="flex flex-1 w-2/3 mx-auto mt-12 border border-solid border-gray-300 p-4">
                <div class="w-1/2">
                    <h2>Login</h2>
                    <div class="flex flex-1">
                        <div class="w-1/2">
                            <p class="m-0 p-3">username:</p>
                            <p class="m-0 p-3">password:</p>
                        </div>
                        <div class="w-1/2">
                            <input class="p-2 h-4 mt-2 w-4/5 border-0 bg-white outline-none border-b border-solid border-black" type="text" v-model="login_username"><br>
                            <input class="p-2 h-4 mt-2 w-4/5 border-0 bg-white outline-none border-b border-solid border-black" type="password" v-model="login_password">
                            <div class="bg-black p-2 text-white mt-4 w-4/5 cursor-pointer" @click="Login">Login</div>
                        </div>
                    </div>
                </div>
                <div class="w-1/2">
                    <h2>Register</h2>
                    <div class="flex flex-1">
                        <div class="w-1/2">
                            <p class="m-0 p-3">username:</p>
                            <p class="m-0 p-3">email:</p>
                            <p class="m-0 p-3">password:</p>
                        </div>
                        <div class="w-1/2">
                            <input class="p-2 h-4 mt-2 w-4/5 border-0 bg-white outline-none border-b border-solid border-black" type="text" v-model="register_username"><br>
                            <input class="p-2 h-4 mt-2 w-4/5 border-0 bg-white outline-none border-b border-solid border-black" type="text" v-model="register_email"><br>
                            <input class="p-2 h-4 mt-2 w-4/5 border-0 bg-white outline-none border-b border-solid border-black" type="password" v-model="register_password">
                            <div class="bg-black p-2 text-white mt-4 w-4/5 cursor-pointer" @click="Register">Register</div>
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
    name: 'Auth',

    data() {
        return {
            login_username: '',
            login_password: '',

            register_username: '',
            register_password: '',
            register_email: '',

            err: ''
        }
    },

    /*
        Auth calls:
        Register: /auth/register POST
            requires:
                username: min:3
                password: min:3
                email
        Login: /auth/login POST
            requires:
                username: min:3
                password: min:3
        Logout: /auth/logout GET

        {success: false/true}
    */

   components: {
    'navbar': Navbar,
    'cart': Cart
  },

   methods: {
       Login() {
           axios.post('/api/auth/login', {
               username: this.login_username,
               password: this.login_password
           }).then(response => {
               if(response.data.success) {
                   this.$router.push('/')
               }else{
                   this.err = "Could not log in!"
               }
           })
       },

       Register() {
           axios.post('/api/auth/register', {
               username: this.register_username,
               password: this.register_password,
               email: this.register_email
           }).then(response => {
               if(response.data.success) {
                   this.$router.push('/')
               }else{
                   this.err = "Could not register!"
               }
           }).catch(err => {
               this.err = "Could not register! " + err
           })
       }
   },
}
</script>
