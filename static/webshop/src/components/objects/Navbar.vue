<template>
    <div id="nav" class="w-full border-b border-solid border-0 border-gray-200">
      <ul class="flex flex-1 list-none m-0 pt-4 pb-4">
        <a href="/products"><li>PRODUCTS</li></a>
        <a href="/categories"><li>CATEGORIES</li></a>
      </ul>
      <div>
        <div :class="this.switch(true)">
          <a href="/auth">Login or Register</a>
        </div>
        <div :class="this.switch(false)">
          <a href="/account">Account</a>
        </div>
      </div>
    </div>
</template>

<script>
export default {

  data() {
    return {
      logged_in: document.cookie.match(/^(.*;)?\s*ses_id\s*=\s*[^;]+(.*)?$/) == null
    }
  },

  methods: {
    switch(insert) {
      return insert == this.logged_in ? "" : "hidden"
    }
  }
}
</script>


<<style>
a {
    @apply .text-black .no-underline;
}
li {
    @apply .ml-2 .mr-2 .p-2 .pl-4 .pr-4 .list-none .text-xs;
}
</style>