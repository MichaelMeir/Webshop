<template>
    <div class="flex flex-1 h-full">
        <div class="w-1/2 h-full overflow-hidden rounded-tl rounded-bl cursor-pointer" @click="openProduct">
            <div class="w-full h-full flex flex-1 justify-center">
                <img class="h-full mx-auto flex-none" :src="this.image_data">
            </div>
        </div>
        <div class="p-2">
            <span class="cursor-pointer" @click="openProduct">{{ product.Name }}</span>
            <div class="text-gray-400">{{ product.Description }}</div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'ProductImage',

    props: ["product", "image_id"],

    data() {
        return {
            image_data: ""
        }
    },

    mounted() {
        axios.get('/api/products/image?id=' + this.image_id).then(response => {
            this.image_data = response.data
        })
    },

    methods: {
        openProduct() {
            this.$router.push('/product/' + this.product.UUID)
        }
    }
}
</script>
