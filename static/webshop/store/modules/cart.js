
import Persisted from 'vuex-persistedstate'
import cookie from 'js-cookie'

const state = {
    products: [],
    total_value: 0
}

const getters = {}

const actions = {}

const mutations = {
    Push(state, product) {
        console.log(product)
        state.total_value += product.Price
        state.products.push(product)
    },

    Remove(state, uuid) {
        for(let i = 0; i < state.products.length; i++) {
            if(state.products[i].UUID == uuid) {
                let out = state.products.splice(i, 1)[0]
                console.log(out)
                state.total_value -= out.Price
                return
            }
        }
    }
}

export default {
    namespaced: true,
    state: state,
    getters: getters,
    actions: actions,
    mutations: mutations,
    plugins: [Persisted({
        storage: {
          getItem: key => cookie.get(key),
          setItem: (key, value) => cookie.set(key, value, { expires: ((1 / 48) * 2) }),
          removeItem: key => cookie.remove(key)
        }
      })]
}