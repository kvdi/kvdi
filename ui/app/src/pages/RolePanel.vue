<!--
Copyright 2020,2021 Avi Zimmerman

This file is part of kvdi.

kvdi is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

kvdi is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with kvdi.  If not, see <https://www.gnu.org/licenses/>.
-->

<template>
  <div class="q-pa-md" stretch>

    <div style="float: right">
      <q-btn flat color="primary" icon-right="add" label="New Role" @click="onNewRole" />
    </div>

    <q-dialog v-model="newRolePrompt" persistent>
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="text-h6">New Role</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-input debounce="500" dense v-model="newRoleName" autofocus :rules="[validateRole]" />
        </q-card-section>

        <q-card-actions align="right" class="text-primary">
          <q-btn flat label="Cancel" v-close-popup />
          <q-btn flat label="Add Role" v-close-popup @click="doCreateRole" />
        </q-card-actions>
      </q-card>
    </q-dialog>

    <div style="clear: right">
      <SkeletonTable v-if="loading"/>

      <q-table
        class="roles-table"
        title="Roles"
        :data="data"
        :columns="columns"
        dense
        row-key="idx"
        v-if="!loading"
        ref="table"
      >
        <template v-slot:body="props">
          <q-tr :props="props">

            <q-td key="name" :props="props">
              <strong>{{ props.row.metadata.name }}</strong>
            </q-td>

            <q-td key="rules" :props="props" text-align="center">
              <div class="rule-wrapper">
                <RuleDisplay
                  v-for="(rule, idx) in props.row.rules"
                  v-bind="rule"
                  :key="'rules' + props.row.idx + '-' + idx"
                  :ref="'rules' + props.row.idx + '-' + idx"
                  :roleIdx="props.row.idx"
                  :roleName="props.row.metadata.name"
                  :ruleIdx="idx"
                  :editable="props.row.editable"
                  style="float: left;"
                />
              </div>
            </q-td>

            <q-td key="annotations" :props="props">
              <RoleAnnotations
                v-bind="props.row"
                :roleIdx="props.row.idx"
                :annotations="props.row.metadata.annotations"
                :editable="props.row.editable"
                :ref="'annotations-' + props.row.idx"
              />
            </q-td>

            <q-td key="updateRole" :props="props">
              <q-btn v-if="props.row.editable" round dense flat icon="add"  size="sm" color="blue" @click="onAddRule(props.row.idx, props.row.metadata.name)">
                <q-tooltip anchor="bottom middle" self="top middle" :offset="[10, 10]">Add Rule</q-tooltip>
              </q-btn>
              &nbsp;
              <q-btn v-if="props.row.editable" round dense flat icon="cancel"  size="sm" color="warning" @click="onCancelEdit(props.row.idx, props.row.metadata.name)">
                <q-tooltip anchor="bottom middle" self="top middle" :offset="[10, 10]">Cancel Editing</q-tooltip>
              </q-btn>
              &nbsp;
              <q-btn v-if="props.row.editable" round dense flat icon="save"  size="sm" color="green" @click="onSaveRole(props.row.idx, props.row.metadata.name)">
                <q-tooltip anchor="bottom middle" self="top middle" :offset="[10, 10]">Save Role</q-tooltip>
              </q-btn>
              <q-btn v-if="!props.row.editable" round dense flat icon="create"  size="sm" color="grey" @click="onEditRole(props.row.idx, props.row.metadata.name)">
                <q-tooltip anchor="bottom middle" self="top middle" :offset="[10, 10]">Edit Role</q-tooltip>
              </q-btn>
              &nbsp;
              <q-btn round dense flat icon="remove_circle"  size="sm" color="red" @click="onConfirmDeleteRole(props.row.idx, props.row.metadata.name)">
                <q-tooltip anchor="bottom middle" self="top middle" :offset="[10, 10]">Delete Role</q-tooltip>
              </q-btn>
            </q-td>

          </q-tr>
        </template>
      </q-table>
    </div>

  </div>
</template>

<script>
import SkeletonTable from 'components/SkeletonTable.vue'

import RuleDisplay from 'components/RuleDisplay.vue'
import RoleAnnotations from 'components/RoleAnnotations.vue'

import ConfirmDelete from 'components/dialogs/ConfirmDelete.vue'

const roleColumns = [
  {
    name: 'name',
    required: true,
    label: 'Role',
    align: 'left',
    field: row => row.name,
    format: val => `${val}`,
    sortable: true,
    classes: 'bg-grey-2 ellipsis',
    headerClasses: 'bg-primary text-white'
  },
  {
    name: 'rules',
    align: 'left',
    label: 'Rules'
  },
  {
    name: 'annotations',
    align: 'left',
    label: 'Annotations'
  },
  {
    name: 'updateRole',
    align: 'center'
  }
]

export default {
  name: 'RolesPanel',
  components: { SkeletonTable, RuleDisplay, RoleAnnotations },

  data () {
    return {
      loading: true,
      data: [],
      columns: roleColumns,
      newRolePrompt: false,
      newRoleName: ''
    }
  },

  created () {
    this.$root.$on('reload-roles', this.fetchData)
  },

  beforeDestroy () {
    this.$root.$off('reload-roles', this.fetchData)
  },

  methods: {

    onNewRole () {
      this.newRolePrompt = true
    },

    onEditRole (roleIdx, roleName) {
      this.$root.$on(roleName, this.doUpdateRole)
      this.data[roleIdx].editable = true
    },

    onSaveRole (roleIdx, roleName) {
      this.$root.$off(roleName, this.doUpdateRole)
      this.data[roleIdx].editable = false
      this.doSaveRole(roleIdx, roleName)
    },

    onAddRule (roleIdx, roleName) {
      if (this.data[roleIdx].rules !== undefined) {
        this.data[roleIdx].rules.push({ verbs: [], resources: [], resourcePatterns: [], namespaces: [] })
        return
      }
      this.data[roleIdx].rules = [{ verbs: [], resources: [], resourcePatterns: [], namespaces: [] }]
    },

    onCancelEdit (roleIdx, roleName) {
      this.$root.$off(roleName, this.doUpdateRole)
      this.data[roleIdx].editable = false
      const annotationRef = this.$refs[`annotations-${roleIdx}`]
      annotationRef.reset()
      this.fetchData()
    },

    onConfirmDeleteRole (roleIdx, roleName) {
      // TODO: There is no server-side check for this yet - and there should be
      if (this.doAdminCheck()) { return }
      this.$q.dialog({
        component: ConfirmDelete,
        parent: this,
        resourceName: roleName
      }).onOk(() => {
        this.doDeleteRole(roleName)
      }).onCancel(() => {
      }).onDismiss(() => {
      })
    },

    doUpdateRole ({ roleIdx, ruleIdx, rulePayload, deleteRule }) {
      if (this.data[roleIdx].rules !== undefined) {
        if (deleteRule === true) {
          this.data[roleIdx].rules.splice(ruleIdx, 1)
          return
        }
        this.data[roleIdx].rules.splice(ruleIdx, 1, rulePayload)
      }
    },

    doAdminCheck (roleName) {
      if (roleName === 'kvdi-admin') {
        this.$q.notify({
          color: 'red-4',
          textColor: 'white',
          icon: 'warning',
          message: 'You cannot delete the kvdi-admin role'
        })
        return true
      }
      return false
    },

    async validateRole (val) {
      if (!val) {
        return 'Name is required'
      }
      try {
        await this.$axios.get(`/api/roles/${val}`)
        return 'Role already exists'
      } catch (err) {}
    },

    async doCreateRole () {
      try {
        await this.$axios.post('/api/roles', { name: this.newRoleName })
        this.$q.notify({
          color: 'green-4',
          textColor: 'white',
          icon: 'cloud_done',
          message: `Created role '${this.newRoleName}'`
        })
      } catch (err) {
        this.$root.$emit('notify-error', err)
      }
      this.fetchData()
    },

    async doSaveRole (roleIdx, roleName) {
      try {
        const annotationRef = this.$refs[`annotations-${roleIdx}`]
        const roleAnnotations = await annotationRef.currentAnnotations()
        const payload = {
          rules: this.data[roleIdx].rules || [],
          annotations: roleAnnotations
        }
        await this.$axios.put(`/api/roles/${roleName}`, payload)
        this.$q.notify({
          color: 'green-4',
          textColor: 'white',
          icon: 'cloud_done',
          message: `Successfully updated role '${roleName}'`
        })
      } catch (err) {
        this.$root.$emit('notify-error', err)
      }
      this.fetchData()
    },

    async doDeleteRole (roleName) {
      try {
        await this.$axios.delete(`/api/roles/${roleName}`)
        this.$q.notify({
          color: 'green-4',
          textColor: 'white',
          icon: 'cloud_done',
          message: `Deleted role '${roleName}'`
        })
      } catch (err) {
        this.$root.$emit('notify-error', err)
      }
      this.fetchData()
    },

    async fetchData () {
      try {
        const res = await this.$axios.get('/api/roles')
        this.data = []
        res.data.forEach((val, idx) => {
          const item = { idx: idx, editable: false, ...val }
          if (item.rules === undefined) {
            item.rules = []
          }
          this.data.push(item)
        })
      } catch (err) {
        this.$root.$emit('notify-error', err)
      }
    }

  },

  async mounted () {
    await this.$nextTick()
    this.loading = true
    await new Promise((resolve, reject) => setTimeout(resolve, 500))
    await this.fetchData()
    this.loading = false
  }
}
</script>

<style lang="sass" scoped>
.rule-wrapper
  position: relative
  width: 100vh

.roles-table
  background-color: $grey-3

  /* height or max-height is important */
  max-height: 500px

  // /* specifying max-width so the example can
  //   highlight the sticky column on any browser window */
  // max-width: 600px

  td:first-child
    /* bg color is important for td; just specify one */
    // background-color: #c1f4cd !important

  tr th
    position: sticky
    /* higher than z-index for td below */
    z-index: 2
    /* bg color is important; just specify one */
    background: #fff

  /* this will be the loading indicator */
  thead tr:last-child th
    /* height of all previous header rows */
    top: 48px
    /* highest z-index */
    z-index: 3
  thead tr:first-child th
    top: 0
    z-index: 1
  tr:first-child th:first-child
    /* highest z-index */
    z-index: 3

  td:first-child
    z-index: 1

  td:first-child, th:first-child
    position: sticky
    left: 0
</style>
