<template lang="pug">
b-container
  b-row
    b-col(sm="4")
      user-nav.mt-2
    b-col(sm="8")
      b-card.mt-2(no-body)
        b-card-body
          h4 WANTED
          p.card-text.text-right
            b-button(variant="outline-primary" v-b-modal.newWanted) new wanted
  b-modal#newWanted(
    title="New Wanted"
    size="lg"
    ref="newWanted")
    b-form
      b-form-group(
        label="Title:"
        label-for="inputTitle")
        b-form-input#inputTitle(
          v-model="title"
          :state="stateTitle"
          placeholder="Please input) title."
          :disabled="modalDisabled")
      b-form-group(
        label="Price(JPY):"
        label-for="inputPrice")
        b-form-input#inputPrice(
          v-model="price"
          type="number"
          placeholder="Please input price."
          :disabled="modalDisabled")
      b-form-group(
        label="Description:"
        label-for=inputDescription)
        b-form-textarea#inputDescription(
          v-model="description"
          :state="stateDescription"
          placeholder="Please input description."
          :rows="6"
          :disabled="modalDisabled")
    div(slot="modal-footer")
      b-button(
        slot="modal-cancel"
        variant="outline-secondary"
        v-on:click="$refs.newWanted.hide()") Cancel
      b-button.ml-2(
        slot="modal-ok"
        variant="outline-primary"
        v-on:click="modalOK"
        :disabled="!stateTitle || !stateDescription") Send
</template>

<script lang="ts">
import {
  Component,
  Vue
} from 'nuxt-property-decorator'
import UserNav from "~/components/UserNav.vue"

@Component({
  components: {
    UserNav,
  }
})
export default class extends Vue {
  public modalDisabled = false;
  
  public title       = "";
  public price       = 10000;
  public description = "";

  public get stateTitle(): boolean {
    return this.title.length > 0 && this.title.length < 100;
  }
  public get stateDescription(): boolean {
    return this.description.length > 0 && this.description.length < 500;
  }

  public modalOK() {
    let modal: any = this.$refs.newWanted;
    modal.hide();
  }
}
</script>

<style lang="sass"></style>
