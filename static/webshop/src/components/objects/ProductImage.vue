<template>
    <img :id="this.new_image" :src="this.image_data">
</template>

<script>
export default {
    name: 'ProductImage',

    props: ['src', 'load'],

    data() {
        return {
            image_data: '',
            image_set: false,
            new_image: this.src
        }
    },

    mounted() {
        this.LoadImage()
    },

    methods: {
        SetImage(src) {
            this.image_set = true
            this.new_image = src
            this.LoadImage()
        },

        LoadImage() {
            if(!this.load) {
                requestIdleCallback(this.LoadImage)
                return
            }
            let url = '/api/products/image?id='
            if(this.image_set) {
                url += this.new_image
            }else{
                url += this.src
            }
            axios.get(url).then(response => {
                this.image_data = response.data
            })
        }
    }
}
</script>

