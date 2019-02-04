<template lang="pug">
  b-card(no-body)
    h4(slot="header")
      b-link(:href="'/user/' + $route.params.id") {{ $route.params.id }}
    b-card-body(v-if="userExist")
      template Rate: ...
    b-card-body(v-else) Not exist.
    b-list-group(v-if="signin" flush)
      b-list-group-item
        b-link(:href="'/works/' + $store.state.name + '/wanted'") My Wanteds
          template(v-if="newWanteds > 0") ({{ newWanteds }})
      b-list-group-item
        b-link(:href="'/works/' + $store.state.name + '/request'") My Requests
          template(v-if="newRequests > 0") ({{ newRequests }})
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import axios from "axios"

@Component
export default class extends Vue {
  public userExist = false;
  public newRequests = 0;
  public newWanteds  = 0;

  public get signin(): boolean {
    return this.userExist && this.$store.state.name === this.$route.params.id;
  }
  private mounted() {
    axios.get("/api/get/user/" + this.$route.params.id).then(({ data }) => {
      if (data.user.name !== "") {
        this.userExist = true;
      }
    });
    axios.get("/api/get/works/info").then(({ data }) => {
      this.newRequests = data.info.requests;
      this.newWanteds  = data.info.wanteds;
    });
  }
}
</script>
