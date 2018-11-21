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

  b-modal#car-modal(size="lg")
    b-row.my-1
      b-col.text-right(sm="1")
        p 車種
      b-col(sm="11")
        b-form-select(v-model="car_name" :options="car_name_options")
    b-row.my-1
      b-col.text-right(sm="1")
        p 年式
      b-col(sm="11")
        b-form-select(v-model="formula_year" :options="formula_year_options")
    b-row.my-1
      b-col.text-right(sm="1")
        p 型式
      b-col(sm="11")
        b-form-input(type="text" v-model="car_model")
    b-row.my-1
      b-col.text-right(sm="1")
        label(for="fuel") 燃料
      b-col
        b-form-radio-group#fuel(v-model="fuel" :options="fuel_options")
    b-row.my-1
      b-col.text-right(sm="1")
        p 乗車定員
      b-col(sm="11")
        b-form-input(type="number" v-model="ridingcapacity")
    b-row.my-1
      b-col.text-right(sm="1")
        p ドア数
      b-col(sm="11")
        b-form-input(type="number" v-model="doors")
    b-row.my-1
      b-col.text-right(sm="1")
        label(for="drivesystem") 駆動方式
      b-col
        b-form-radio-group#fuel(v-model="drivesystem" :options="drivesystem_options")
    b-row.my-1
      b-col.text-right(sm="1")
        label(for="equipment") 装備品
      b-col
        b-form-checkbox-group#equipment(v-model="equipment" :options="equipment_options")


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
      car_name: null,
      formula_year:null,
      car_model: '',
      fuel:null,
      ridingcapacity:0,
      doors:0,
      drivesystem: null,
      equipment:[],

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
      car_name_options: [
        { value: null, text: '車種を選択してください' },
        { value: 1, text: 'bbbオークション' },
        { value: 2, text: 'cccオークション' },
      ],
      formula_year_options: [
        { value: null, text: '年式を選択してください' },
        { value: 1, text: '1997' },
        { value: 2, text: '1998' },
      ],
      fuel_options: [
        {value: '0',  text: 'G'},
        {value: '1',  text: 'D'},
      ],
      drivesystem_options: [
        {value: '0',  text: 'G'},
        {value: '1',  text: 'D'},
      ],
      equipment_options: [
        {value: '0',  text: 'パワステ(PS)'},
        {value: '1',  text: 'パワーウィンドウ(PW)'},
        {value: '2',  text: 'ABS'},
        {value: '3',  text: 'アルミホイール(AW)'},
        {value: '4',  text: 'カセット'},
        {value: '5',  text: 'CD'},
        {value: '6',  text: 'MD'},
        {value: '7',  text: 'TV'},
        {value: '8',  text: 'ナビ'},
        {value: '9',  text: 'エアバック(シングル)'},
        {value: '10',  text: 'エアバック(W)'},
        {value: '11',  text: 'TVリモコン'},
        {value: '12',  text: 'ナビリモコン'},
        {value: '13',  text: 'エアリモコン'},
        {value: '14',  text: 'オーディオリモコン'},
        {value: '15',  text: 'エアコン'},
      ],
    }
  },
  mounted() {
    this.id = this.$route.params.id
    this.info_header.next.url = `/shipment_info/new?purchase_id=${this.id}`
  }
}
</script>
