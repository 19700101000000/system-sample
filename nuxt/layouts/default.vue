<template lang="pug">
div
  b-navbar(toggleable="md" type="dark" variant="info")
    b-navbar-toggle(target="nav_collapse")
    b-navbar-brand(href="/") IH-SYSTEM
    b-collapse#nav_collapse(is-nav)
      b-navbar-nav
        b-nav-item(href="/buy") 買注文
        b-nav-item(href="/sell") 売注文
      b-navbar-nav.ml-auto
        b-nav-item-dropdown(
          v-if="$store.state.auth"
          :text="$store.state.auth.username"
          right
        )
          b-dropdown-item(href="#") プロフィール
          b-dropdown-item(href="#" @click="$store.dispatch('logout')") ログアウト
        b-nav-item(href="/login" v-else right) ログイン
  b-container
    nuxt
</template>

