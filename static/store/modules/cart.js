import { stat } from "fs";


const state = {
    products: [],
    total_value: 0
}

const getters = {

}

const actions = {
    AddProduct({ state, commit }, product) {
        commit('Push', {product: product})
    },

    RemoveProduct({ state, commit }, product) {
        commit('Remove', {id: product.UUID})
    }
}

const mutations = {
    Push(state, { product }) {
        state.total_value += product.Price
        state.products.push(product)
    },

    Remove(state, { id }) {
        for(let i = 0; i < state.products.length; i++) {
            if(state.products[i].UUID = id) {
                state.products.splice(i, 1)
                return
            }
        }
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}