<template>
    <div>
        {{ product.Name }}
        <span class="text-gray-400">{{ product.Description }}</span>
        <img :src="image_data">
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
        axios.get('http://localhost/api/products/image?id=' + this.image_id).then(response => {
            this.image_data = response.data
        })
    }
}
</script>
