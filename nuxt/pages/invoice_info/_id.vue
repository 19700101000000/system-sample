<template lang="pug">
div
  info-header(:info="info_header")
  b-row.my-1
    b-col
      b-card(title="請求")
        b-row.my-1
          b-col.text-right(sm="1")
            p NO
          b-col(sm="11")
            p {{id}}

        b-row.my-1(align-h="center")
          b-col.text-right(sm="1")
            label(for="client_name") 顧客
          b-col(sm="5")
            b-form-select#client_name(type="text" v-model="client_name" :options="client_name_options")
          b-col.text-right(sm="1")
            label(for="employee_name") 担当者
          b-col(sm="5")
            b-form-select#employee_name(type="text" v-model="employee_name" :options="employee_name_options")

  b-row.my-1(align-h="center")
    b-col(sm="2")
      b-button(variant="outline-success" block) 登録
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
  data(){
    return{
      info_header: {
        next: {
          text: '回収',
          url: '',
        }
      },
      id: '',
      client_name  :null,
      employee_name:null,

      client_name_options: [
        { value: null, text: '顧客を選択してください' },
        { value: 'a', text: '高田1世' },
        { value: 'b', text: '高田2世' },
        { value: 'c', text: '高田3世' },
      ],
      employee_name_options: [
        { value: null, text: '担当者を選択してください' },
        { value: 'a', text: 'ぶちお' },
        { value: 'b', text: 'ぶちこ' },
        { value: 'c', text: 'プー太郎' },
      ]
    }
  },
  mounted() {
    this.id = this.$route.params.id
    this.info_header.next.url = `/recovery_info/new?invoice_id=${this.id}`
  }
}
</script>
