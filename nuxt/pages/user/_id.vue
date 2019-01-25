<template lang="pug">
b-container
  b-row
    b-col(sm="4")
      b-card.mt-2(no-body)
        h4(slot="header") {{ $route.params.id }}
        b-card-body(v-if="userExist")
          template Rate: ...
        b-card-body(v-else) Not exist.
        b-list-group(v-if="signin" flush)
          b-list-group-item
            b-link 仕事依頼
          b-list-group-item
            b-link 仕事募集
    b-col(sm="8")
</template>

<script lang="ts">
import {
  Component,
  Vue
} from 'nuxt-property-decorator'
import axios from "axios"

@Component({})
export default class extends Vue {
  public userExist = false;

  public get signin(): boolean {
    return this.userExist && this.$store.state.name === this.$route.params.id;
  }
  private mounted() {
    axios.get("/api/get/user/" + this.$route.params.id).then(({ data }) => {
      if (data.user.name !== "") {
        this.userExist = true;
      }
    });
  }
}
</script>

<style lang="sass"></style>
