<template>
    <q-page padding>
        <div class="row">
            <div class="col-4 column items-center q-gutter-md">
                <q-date minimal v-model="queryParams.inAreaTime" mask="YYYY-MM-DD" @update:model-value="changeDate" />
                <q-card style="width: 80%; max-width: 400px; padding: 10px" v-if="dayList.length">
                    <q-card-section horizontal>
                        <GqaAvatar size="150px" />
                        <q-card-section class="column items-center justify-center" style="width: 100%">
                            <q-chip class="glossy" color="orange" text-color="white" icon-right="star">
                                今日最先打卡
                            </q-chip>

                            <span>
                                {{ showDateTime(dayList[0].inAreaTime) }}

                            </span>
                            <span>
                                {{ dayList[0].userName }}(
                                {{ dayList[0].workNumber }}
                                )

                            </span>
                        </q-card-section>
                    </q-card-section>
                </q-card>
                <q-card style="width: 80%; max-width: 400px; padding: 10px" v-if="yearFirstList.length">
                    <q-card-section horizontal>
                        <GqaAvatar size="150px" />
                        <q-card-section class="column items-center justify-center" style="width: 100%">
                            <q-chip class="glossy" color="deep-orange" text-color="white" icon-right="star">
                                本年度最先打卡
                            </q-chip>

                            <span>
                                {{ showDateTime(yearFirstList[0].inAreaTime) }}

                            </span>
                            <span>
                                {{ yearFirstList[0].userName }}(
                                {{ yearFirstList[0].workNumber }}
                                )

                            </span>
                        </q-card-section>
                    </q-card-section>
                </q-card>
            </div>
            <div class="col">
                <div class="items-center row q-gutter-md" style="margin-bottom: 10px">
                    <q-input style="width: 20%" v-model="queryParams.workNumber" label="工号" />
                    <q-input style="width: 20%" v-model="queryParams.userName" label="姓名" />
                    <q-btn color="primary" @click="handleSearch" :label="$t('Search')" />
                    <q-btn color="primary" @click="resetSearch" :label="$t('Reset')" />
                </div>
                <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns"
                    v-model:pagination="pagination" :rows-per-page-options="pageOptions" :loading="loading"
                    @request="onRequest">

                    <template v-slot:body-cell-inAreaTime="props">
                        <q-td :props="props">
                            {{ showDateTime(props.row.inAreaTime) }}
                        </q-td>
                    </template>

                    <template v-slot:body-cell-actions="props">
                        <q-td :props="props">
                            <div class="q-gutter-xs">
                                <q-btn dense color="primary" @click="personResult(props.row)" label="个人记录" />
                            </div>
                        </q-td>
                    </template>
                </q-table>
            </div>
        </div>

        <q-dialog v-model="personVisible">
            <q-spinner color="primary" size="3em" v-if="personLoading" />
            <q-card style="width: 100%; max-width: 400px; padding: 10px" v-if="personFirstList.length">
                <q-card-section horizontal>
                    <GqaAvatar size="150px" />
                    <q-card-section class="column items-center justify-center" style="width: 100%">
                        <span>
                            {{ personFirstList[0].userName }}(
                            {{ personFirstList[0].workNumber }}
                            )
                        </span>
                        <q-chip class="glossy" color="deep-orange" text-color="white" icon-right="star">
                            本年度最先打卡
                        </q-chip>
                        <span>
                            创造于：
                        </span>
                        <span>
                            {{ showDateTime(personFirstList[0].inAreaTime) }}
                        </span>
                    </q-card-section>
                </q-card-section>
            </q-card>

        </q-dialog>

    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { postAction } from 'src/api/manage'
import { date } from 'quasar'
import { computed, ref, onBeforeMount } from 'vue'

const url = {
    list: 'plugin-attendance/in-area-list',
    year: 'plugin-attendance/in-area-year',
}
const columns = [
    { name: 'workNumber', align: 'center', label: '工号', field: 'workNumber' },
    { name: 'userName', align: 'center', label: '姓名', field: 'userName' },
    { name: 'inAreaTime', align: 'center', label: '打卡时间', field: 'inAreaTime' },
    { name: 'actions', align: 'center', label: '操作', field: 'actions' },
]
const {
    showDateTime,
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
    GqaAvatar,
    loading,
    tableData,
    recordDetailDialog,
    showAddForm,
    showEditForm,
    onRequest,
    getTableData,
    handleSearch,
    // resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

const dayList = ref([])
const yearFirstList = ref([])
const personVisible = ref(false)
const personFirstList = ref([])
const personLoading = ref(false)

onBeforeMount(() => {
    const dateNow = new Date()
    queryParams.value.inAreaTime = date.formatDate(dateNow, 'YYYY-MM-DD')
    pagination.value = {
        sortBy: 'created_at',
        descending: true,
        page: 1,
        rowsPerPage: 10,
    }
    getDayResult()
    getYearResult()
})
const getDayResult = async () => {
    dayList.value = []
    await getTableData()
    dayList.value = tableData.value
}
const getYearResult = () => {
    yearFirstList.value = []
    postAction(url.year, {
        workNumber: '',
    }).then((res) => {
        if (res.code === 1) {
            yearFirstList.value = res.data.records
        }
    })
}
const changeDate = (value) => {
    pagination.value = {
        sortBy: 'created_at',
        descending: true,
        page: 1,
        rowsPerPage: 10,
    }
    queryParams.value.inAreaTime = value
    getDayResult()
}
const personResult = (row) => {
    personFirstList.value = []
    personLoading.value = true
    personVisible.value = true
    postAction(url.year, {
        workNumber: row.workNumber,
    }).then((res) => {
        if (res.code === 1) {
            personFirstList.value = res.data.records
        }
    }).finally(() => {
        personLoading.value = false
    })
}
const resetSearch = () => {
    queryParams.value = {}
    const dateNow = new Date()
    pagination.value = {
        sortBy: 'created_at',
        descending: true,
        page: 1,
        rowsPerPage: 10,
    }
    queryParams.value.inAreaTime = date.formatDate(dateNow, 'YYYY-MM-DD')
    getTableData()
}
</script>
