<template lang="pug">
b-container
  b-row
    b-col(sm="4")
      user-nav.mt-2
    b-col(sm="8")
      b-card.mt-2(no-body)
        b-card-body
          h4 MY REQUESTS
      request-list-item.mt-2(v-for="request in requests" :value="request" :wanted="true")
</template>

<script lang="ts">
import {
  Component,
  Vue
} from 'nuxt-property-decorator'
import axios from "axios"
import UserNav from "~/components/UserNav.vue"
import RequestListItem from "~/components/RequestListItem.vue"

@Component({
  components: {
    UserNav,
    RequestListItem,
  }
})
export default class extends Vue {
  public requests = []
  public mounted() {
    axios.get("/api/get/works/requests").then(({ data }) => {
      this.requests = data.requests;
    });
  }
}
</script>

<style lang="sass"></style>
