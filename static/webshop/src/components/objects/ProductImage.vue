<template>
    <img :src="this.image_data">
</template>

<script>
export default {
    name: 'ProductImage',

    props: ['src', 'load'],

    data() {
        return {
            image_data: ''
        }
    },

    mounted() {
        this.LoadImage()
    },

    methods: {
        LoadImage() {
            if(!this.load) {
                requestIdleCallback(this.LoadImage)
                return
            }
            axios.get('http://localhost/api/products/image?id=' + this.src).then(response => {
                this.image_data = response.data
            })
        }
    }
}
</script>

