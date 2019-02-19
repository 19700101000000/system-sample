<template lang="pug">
b-card(no-body)
  h4(slot="header")
    b-link(:href="'/user/' + $route.params.id") {{ $route.params.id }}
  b-card-body(v-if="userExist")
    b-link(v-on:click="showRateModal") Rate: {{ userRate }}
    b-link.ml-2(v-if="user.name !== $store.state.name && user.requests > 0" v-b-modal.evaluateModal) evalutate to {{ user.name }}
  b-card-body(v-else) Not exist.
  b-list-group(flush)
    b-list-group-item
      b-link(:href="'/gallery/' + $route.params.id")
        template(v-if="signin") My Galleries
        template(v-else) {{ $route.params.id }}'s Galleries
    template(v-if="signin")
      b-list-group-item
        b-link(:href="'/works/' + $store.state.name + '/wanted'") My Wanteds
          template(v-if="newWanteds > 0") ({{ newWanteds }})
      b-list-group-item
        b-link(:href="'/works/' + $store.state.name + '/request'") My Requests
          template(v-if="newRequests > 0") ({{ newRequests }})
  b-card-body
  b-modal#evaluateModal(
    :title="'Evaluate to ' + $store.state.name"
    ref="modalEval"
    size="lg")
    b-form
      b-form-group(
        label="Rate:"
        label-for="inputRate")
        b-form-input#inputRate(
          v-model="evaluate.rate"
          :formatter="rateFormat"
          :state="stateRate"
          placeholder="Please input Rate(0 ~ 5).")
      b-form-group(
        label="Review:"
        label-for="inputReview")
        b-form-textarea#inputReview(
          v-model="evaluate.review"
          :state="stateReview"
          :rows="6"
          placeholder="Please input Review.")
        p.text-right {{ evaluate.review.length}} / 500
    p.text-danger(v-if="error") Fail connection.
    div(slot="modal-footer")
      b-button.mr-2(variant="outline-secondary" v-on:click="$refs.modalEval.hide()") Cancel
      b-button(variant="outline-primary" :disabled="!stateRate || !stateReview" v-on:click="onSendEval") Send
  modal-rank(ref="modalRank" :userName="$route.params.id")
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import axios from "axios"
import ModalRank from "~/components/ModalRank.vue"

interface user {
  name: string,
  showname: string,
  alive: boolean,
  rate: number | null,
  base: number | null,
  monitor: {
    rate: number | null,
    review: string | null,
  },
  requests: number,
}

@Component({
  components: {
    ModalRank,
  }
})
export default class extends Vue {
  public userExist = false;
  public newRequests = 0;
  public newWanteds  = 0;

  public error =false;

  public user: user;

  public evaluate = {
    rate: "0",
    rateTrue: 0,
    review: "foo",
  };

  public get userRate(): number {
    if (this.user.rate && this.user.base && this.user.base > 0) {
      return Math.round(this.user.rate / this.user.base * 10) / 10
    }
    return 0
  }

  public get stateRate(): boolean {
    return this.evaluate.rateTrue > 0 && this.evaluate.rateTrue <= 5 && this.evaluate.rate.length === 1;
  }
  public get stateReview(): boolean {
    return this.evaluate.review.length > 0 && this.evaluate.review.length <= 500;
  }

  public rateFormat(value: string, event: Event): string {
    const match = value.match(/[0-5]/)
    if (match) {
      this.evaluate.rateTrue = parseInt(match[0]);
      return match[0];
    }
    return "";
  }

  public get signin(): boolean {
    return this.userExist && this.$store.state.name === this.$route.params.id;
  }

  public onSendEval(event: Event) {
    const modal: any = this.$refs.modalEval;
    axios.post("/api/update/evaluate", {
      user: this.user.name,
      rate: this.evaluate.rateTrue,
      review: this.evaluate.review,
    }).then(({ data }) => {
      this.error = false;
      modal.hide()
    }).catch((result) => {
      this.error = true;
    });
  }

  public showRateModal() {
    const modal: any = this.$refs.modalRank;
    modal.showModal();
  }

  private mounted() {
    axios.get("/api/get/user/" + this.$route.params.id).then(({ data }) => {
      this.user = data.user;
      if (this.user.name !== "") {
        this.userExist = true;
      }
      if (this.user.monitor.rate && this.user.monitor.review) {
        this.evaluate.rate = this.user.monitor.rate.toString();
        this.evaluate.review = this.user.monitor.review;
      }
    });
    axios.get("/api/get/works/info").then(({ data }) => {
      this.newRequests = data.info.requests;
      this.newWanteds  = data.info.wanteds;
    });
  }
}
</script>
