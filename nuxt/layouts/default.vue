<template lang="pug">
  div
    b-navbar(toggleable="md", type="dark", variant="info")
      b-navbar-toggle(target="nav_collapse")
      b-navbar-brand(href="#") Top
      b-collapse#nav_collapse(is-nav)
        b-navbar-nav
          b-nav-item(href="#") New
          b-nav-item(href="#") Hot
          b-nav-item(href="#") Pickup
        b-navbar-nav.ml-auto
          b-nav-item(v-if="false" href="#" right) Signin
          b-nav-item-dropdown(v-else text="User" right)
            b-dropdown-item(href="#") Profile
            b-dropdown-item(href="#") Signout
    nuxt
</template>
