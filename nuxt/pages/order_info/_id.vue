<template lang="pug">
div
  info-header(:info="info_header")
  b-row.my-1
    b-col
      b-card.my-1(:title="page_items.title")
        b-row.my-1
          b-col.text-right(sm="1")
            p NO
          b-col(sm="11")
            p {{id}}

        b-row.my-1
          b-col.text-right(sm="1")
            label(for="client_name") 顧客
          b-col(sm="5")
            b-form-select#client_name(type="text" v-model="client_name" :options="client_name_options")
          b-col.text-right(sm="1")
            label(for="employee_name") 担当者
          b-col(sm="5")
            b-form-select#employee_name(type="text" v-model="employee_name" :options="employee_name_options")

        b-row.my-1
          b-col.text-right(sm="1")
            label(for="car_budget") 予算
          b-col
            b-form-input#car_budget(type="number" v-model="car_budget" placeholder="0")

        b-row.my-1
          b-col.text-right(sm="1")
            label(for="manufacturer") メーカー
          b-col
            b-form-select#manufacturer(v-model="manufacturer" :options="manufacturer_options")

        b-row.my-1
          b-col.text-right(sm="1")
            label(for="car_name") 車種
          b-col
            b-form-select#car_name(v-model="car_name" :options="car_name_options")

        b-row.my-1
          b-col.text-right(sm="1")
            label(for="car_model_year") 年式
          b-col(sm="3")
            b-form-select#car_model_year(v-model="car_model_jpc_name" :options="car_model_jpc_name_options")
          b-col(sm="8")
            b-form-input#car_model_year(type="number" v-model="car_model_jpc_year")

        b-row.my-1
          b-col.text-right(sm="1")
            label(for="car_model_no") 型式
          b-col
            b-form-select#car_model_no(v-model="car_model_no" :options="car_model_no_options")

        b-row.my-1
          b-col.text-right(sm="1")
            label(for="car_exterior_color") 外装色
          b-col
            b-form-checkbox-group#car_exterior_color(v-model="car_exterior_color" :options="car_exterior_color_options")

        b-row.my-1
          b-col.text-right(sm="1")
            label(for="car_interior_color") 内装色
          b-col
            b-form-checkbox-group#car_interior_color(v-model="car_interior_color" :options="car_interior_color_options")

  b-row.my-1
    b-col(sm="12")
      b-progress(:value="page_items.progress.value" :max="page_items.progress.max" animated v-if="page_items.progress.show")
      b-alert(variant="danger" dismissible :show="page_items.danger_alert.show" @dismissed="dismissed") {{ page_items.danger_alert.msg }}
  b-row.my-1(align-h="center")
    b-col(sm="2")
      b-button(variant="outline-success" block @click="onClick" v-if="page_items.update_button.show") {{ page_items.update_button.msg }}

</template>

<style lang="sass" scope>
</style>

<script>
import InfoHeader from '~/components/InfoHeader'
import axios from 'axios'
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
  data() {
    return {
      page_items: {
        title: '受注',
        progress: {
          show: false,
          value: 0,
          max: 5,
        },
        update_button: {
          show: true,
          msg: '登録',
        },
        danger_alert: {
          show: false,
          msg: '登録に失敗しました',
        }
      },
      info_header: {
        next: {
          url: '',
          text: '仕入',
        },
      },
      id: '',
      client_name  :null,
      employee_name:null,
      car_budget   :  '',
      manufacturer :  null,
      car_name     :  null,
      car_model_year: null,
      car_model_no:   null,
      car_model_jpc_name: null,
      car_model_jpc_year: null,
      car_exterior_color: [],
      car_interior_color: [],
      estimate_flag: true,
      new_order: true,

      client_name_options: [
        { value: null, text: '顧客を選択してください' },
      ],
      employee_name_options: [
        { value: null, text: '担当者を選択してください' },
      ],
      manufacturer_options: [
        { value: null, text: 'メーカを選択してください' },
        { value: 'a', text: 'ホンダ' },
        { value: 'b', text: 'スバル' },
        { value: 'c', text: 'トヨタ' },
      ],
      car_name_options: [
        { value: null, text: '車種を選択してください' },
        { value: 'a', text: 'ホンダ' },
        { value: 'b', text: 'スバル' },
        { value: 'c', text: 'トヨタ' },
      ],
      car_model_jpc_name_options: [
        { value: null, text: '年式を選択してください' },
        { value: 'a', text: '昭和' },
        { value: 'b', text: '平成' },
      ],
      car_model_no_options:[
        { value: null, text: '型式を選択してください' },
        { value: 'a', text: '802' },
        { value: 'b', text: '320' },
        { value: 'c', text: '408' },
      ],
      car_exterior_color_options: [],
      car_interior_color_options: []
    }
  },
  methods: {
    onClick() {
      this.page_items.update_button.show = false
      this.page_items.progress.show = true
      /* 画面更新待ちに使用 */
      this.$nextTick(async function() {
        this.page_items.progress.value = this.page_items.progress.max - 1

        try {
          const { data } = await axios.post('/api/update/order')
          if (data.result) { /* success */
            this.page_items.danger_alert.show = true
          } else { /* fail */
          }
        } catch(error) {
          console.log(error.message)
        }
        this.page_items.progress.value = this.page_items.progress.max
        this.page_items.progress.show = false
      })
    },
    dismissed() {
      this.page_items.danger_alert.show = false
      this.page_items.update_button.show = true
      this.page_items.progress.value = 0
    },
    getEmployees: async function() {
      try {
        const { data } = await axios.post('/api/item/employees')
        this.employee_name_options = this.employee_name_options.concat(data)
      } catch(error) {
        console.log(error.message)
      }
    },
    getClients: async function() {
      try {
        const { data } = await axios.post('/api/item/clients')
        this.client_name_options = this.client_name_options.concat(data)
      } catch(error) {
        console.log(error.message)
      }
    },
    getColors: async function() {
      try {
        const { data } = await axios.post('/api/item/colors')
        this.car_exterior_color_options = data
        this.car_interior_color_options = data
      } catch(error) {
        console.log(error.message)
      }
    },
  },
  mounted() {
    this.getEmployees()
    this.getClients()
    this.getColors()

    this.id = this.$route.params.id
    if(this.id === 'new' && this.$route.query.estimate_id) {
      /* new 受注 */
      this.page_items.title = '新規受注'
      this.info_header.next = null
      this.estimate_flag = false
      this.new_order = true
      this.id = this.$route.query.estimate_id
    } else if(this.id === 'new') {
      /* new 見積 */
      this.page_items.title = '新規見積'
      this.info_header.next = null
      this.estimate_flag = true
      this.new_order = true
    } else {
      this.info_header.next.url = `/purchase_info/new?order_id=${this.id}`
      // this.$route.query.order_id
      this.new_order = false
      this.page_items.update_button.msg = '更新'
    }
  }
}
</script>
