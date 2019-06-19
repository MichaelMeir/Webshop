<template>
    <div id="nav" class="w-full border-b border-solid border-0 border-gray-200">
      <ul class="flex flex-1 list-none m-0 pt-4 pb-4">
        <a href="/products"><li>PRODUCTS</li></a>
        <a href="/categories"><li>CATEGORIES</li></a>
        <li>
          <div :class="this.switch(true)">
            <a href="/auth">Login or Register</a>
          </div>
          <div :class="'flex-1 flex ' + this.switch(false)">
            <a href="/account" class="mr-2">Account</a>
            <div @click="logout()" class="cursor-pointer">
              logout
            </div>
          </div>
        </li>
      </ul>
    </div>
</template>

<script>
export default {

  data() {
    return {
      logged_in: (this.getCookie('ses_id') == null || this.getCookie('ses_id').length == 0)
    }
  },

  methods: {
    getCookie(name) {
      var value = "; " + document.cookie;
      var parts = value.split("; " + name + "=");
      if (parts.length == 2) return parts.pop().split(";").shift();
    },

    switch(insert) {
      return insert == this.logged_in ? "" : "hidden"
    },

    logout() {
      axios.get('/api/auth/logout').then(response => {
        this.$router.push('/')
        this.logged_in = (this.getCookie('ses_id') == null || this.getCookie('ses_id').length == 0)
      })
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