<template lang="pug">
b-container
  b-row
    b-col(sm="4")
      user-nav.mt-2
    b-col(sm="8" v-if="$store.state.name !== ''")
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
        :label="'Title ('+ countTitle  +'):'"
        label-for="inputTitle")
        b-form-input#inputTitle(
          v-model="title"
          :state="stateTitle"
          placeholder="Please input) title."
          :disabled="modalDisabled")
      b-form-group(
        :label="'Description ('+ countDescription +'):'"
        label-for=inputDescription)
        b-form-textarea#inputDescription(
          v-model="description"
          :state="stateDescription"
          placeholder="Please input description."
          :rows="6"
          :disabled="modalDisabled")
      b-form-group(
        label="Price(JPY):"
        label-for="inputPrice")
        b-form-input#inputPrice(
          :class="[price.length > 0 ? 'text-right' : '']"
          v-model="price"
          :state="statePrice"
          :formatter="priceFormat"
          placeholder="Please input price."
          :disabled="modalDisabled")
    div.text-danger(v-if="error") {{ errorMsg }}
    div(slot="modal-footer")
      b-button(
        slot="modal-cancel"
        variant="outline-secondary"
        v-on:click="$refs.newWanted.hide()") Cancel
      b-button.ml-2(
        slot="modal-ok"
        variant="outline-primary"
        v-on:click="modalOK"
        :disabled="!stateTitle || !stateDescription || !statePrice || modalDisabled") Send
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import UserNav from "~/components/UserNav.vue"
import axios from "axios"

@Component({
  components: {
    UserNav,
  }
})
export default class extends Vue {
  public modalDisabled = false;
  
  public error = false;
  public errorMsg = "Fail Uploaded.";
  
  public title       = "";
  public price       = "10,000";
  public priceTrue   = 0;
  public description = "";

  public get countTitle(): number {
    return 100 - this.title.length;
  }
  public get countDescription(): number {
    return 500 - this.description.length;
  }
  public priceFormat(value: string, event): string {
    const match = value.replace(/,/g, '').match(/\d+/);
    if (match) {
      const price = match[0].replace(/(^0+)(\d+)/, '$2');
      this.priceTrue = parseInt(price);

      return price.replace( /(\d)(?=(\d\d\d)+(?!\d))/g, "$1,");
    }
    return "";
  }

  public get stateTitle(): boolean {
    return this.title.length > 0 && this.title.length <= 100;
  }
  public get stateDescription(): boolean {
    return this.description.length > 0 && this.description.length <= 500;
  }
  public get statePrice(): boolean {
    return this.price.length > 0;
  }

  public modalOK() {
    this.modalDisabled = true;
    axios.post("/api/upload/wanted", {
      title: this.title,
      description: this.description,
      price: this.priceTrue,
    }).then((result) => {
      this.modalDisabled = false;
      this.title = "";
      this.description = "";
      this.price = "10000";
      this.priceTrue = 10000;
      this.error = false;
      let modal: any = this.$refs.newWanted;
      modal.hide();
    }).catch((result) => {
      this.error = true;
      this.modalDisabled = false;
    });
  }
}
</script>

<style lang="sass"></style>
