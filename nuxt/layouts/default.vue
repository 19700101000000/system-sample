<template lang="pug">
  div
    b-navbar(toggleable="md", type="dark", variant="dark")
      b-navbar-toggle(target="nav_collapse")
      b-navbar-brand(href="/") Top
      b-collapse#nav_collapse(is-nav)

        b-navbar-nav.ml-auto(v-if="isSignin")
          b-nav-item-dropdown(:text="userName" right)
            b-dropdown-item(href="/user/1") Profile
            b-dropdown-item(v-on:click="signout") Signout

        b-navbar-nav.ml-auto(v-else)
          b-nav-item(href="/signin" right) Sign in
          b-button(href="/signup" variant="outline-info") Sign up
    nuxt
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import axios from "axios"

@Component
export default class extends Vue {
  public userName = "user";
  public isSignin: boolean = false;

  public signout(): void {
    let url = "/api/auth/signout";
    axios.get(url)
    .then((result) => {
      let status = result.data.status;
      if (status === "ok") {
        window.document.location.href = "/";
      }
    });
  }
  public mounted() {
    let url = "/api/auth/check";
    axios.get(url)
    .then((result) => {
      let status = result.data.status;
      if (status === "ok") {
        this.isSignin = true
        this.userName = result.data.name;
      }
    });
  }
}
</script>
