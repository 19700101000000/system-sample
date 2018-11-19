<template lang="pug">
div
  info-header(:info="info_header")
  b-row.my-1
    b-col
      b-card(title="仕入")
        b-row.my-1
          b-col.text-right(sm="1")
            p NO
          b-col(sm="11")
            p {{id}}

        b-row.my-1
          b-col.text-right(sm="1")
            p 顧客
          b-col(sm="5")
            b-form-select(v-model="client" :options="client_options")
          b-col.text-right(sm="1")
            p 担当者
          b-col(sm="5")
            b-form-select(v-model="employee" :options="employee_options")

        b-row.my-1
          b-col.text-right(sm="1")
            p 仕入日
          b-col(sm="11")
            b-form-input(type="date" v-model="purchase_date")

        b-row.my-1
          b-col.text-right(sm="1")
            p 仕入先
          b-col(sm="11")
            b-form-select(v-model="supplier" :options="supplier_options")

        b-row.my-1
          b-col.text-right(sm="1")
            p 出品番号
          b-col(sm="11")
            b-form-input(type="text" v-model="listing_no")

        b-row.my-1
          b-col.text-right(sm="1")
            p 名変期限
          b-col(sm="11")
            b-form-input(type="date" v-model="car_name_deadline")

        b-row.my-1
          b-col.text-right(sm="1")
            p 仕入原価
          b-col(sm="11")
            b-form-input(type="number" v-model="purchase_cost")

        b-row.my-1
          b-col.text-right(sm="1")
            p 手数料
          b-col(sm="11")
            b-form-input(type="number" v-model="fee")

        b-row.my-1
          b-col.text-right(sm="1")
            p 車
          b-col(sm="11")
            b-button(variant="outline-primary" v-b-modal.car-modal block) 車詳細

  b-row.my-1(align-h="center")
    b-col(sm="2")
      b-button(variant="outline-success" block) 登録

  b-modal#car-modal ここに車の詳細を記載する

</template>

<script>
import InfoHeader from '~/components/InfoHeader'
export default {
  validate({params}) {
    return /^\d+$|^new$/.test(params.id)
  },
  fetch({store, redirect}) {
    if (!store.state.auth) {
      return redirect('/login')
    }
  },
  components: {
    InfoHeader
  },
  data () {
    return {
      info_header: {
        next: {
          url: '',
          text: '出荷',
        }
      },
      id: '',
      client: null,
      employee: null,
      purchase_date: '1970-01-01',
      supplier: null,
      listing_no: '',
      car_name_deadline: '1970-01-01',
      purchase_cost: 0,
      fee: 100,

      client_options: [
        { value: null, text: 'aaa' },
        { value: 1, text: 'bbb' },
        { value: 2, text: 'ccc' },
      ],
      employee_options: [
        { value: null, text: 'AAA' },
        { value: 1, text: 'BBB' },
        { value: 2, text: 'CCC' },
      ],
      supplier_options: [
        { value: null, text: '仕入先を選択してください' },
        { value: 1, text: 'bbbオークション' },
        { value: 2, text: 'cccオークション' },
      ],
    }
  },
  mounted() {
    this.id = this.$route.params.id
    this.info_header.next.url = `/shipment_info/new?purchase_id=${this.id}`
  }
}
</script>
