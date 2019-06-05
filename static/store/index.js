import Vue from 'vue'
import Vuex from 'vuex'

import Cart from './modules/cart'

Vue.use(Vuex)

const store = new Vuex.Store({
    modules: {
        Cart
    }
})