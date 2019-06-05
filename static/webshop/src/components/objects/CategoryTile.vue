<template>
    <div @click="openCategory" class="bg-black text-white cursor-pointer p-4">
        {{ this.category.Name }}
    </div>
</template>

<script>
export default {
    name: 'CategoryTile',

    props: ["category"],

    methods: {
        openCategory() {
            this.$router.push('/category/' + this.category.Name)
        }
    }
}
</script>
