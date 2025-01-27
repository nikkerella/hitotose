<template>
  <div>
    <v-container>
      <v-card>
        <v-layout>
          <v-app-bar color="pink-darken-3" image="">
            <template v-slot:image>
              <v-img gradient="to top right, rgba(19,84,122,0.4), rgba(128,208,199,0.3)"></v-img>
            </template>

            <v-toolbar-title>랭킹</v-toolbar-title>

            <v-spacer></v-spacer>

            <v-btn v-if="token" icon class="ml-auto" @click="createRankDialog = true">
              <v-icon>mdi-database-plus-outline</v-icon>
            </v-btn>

          </v-app-bar>

          <v-main>
            <v-container>
              <v-table fixed-header>
                <thead>
                  <tr>
                    <th class="text-left">Title</th>
                    <th class="text-left">Members</th>
                    <th class="text-left">1st</th>
                    <th class="text-left">2nd</th>
                    <th class="text-left">3rd</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="rank in ranks" :key="rank.id" @click="handleRanksRowClick(rank.id)">
                    <td>
                      <span v-if="token" @click="handleShowRankMembersDialog(rank.id)"
                        style="cursor: pointer; text-decoration: underline;">
                        {{ rank.title }}
                      </span>
                      <span v-else>
                        {{ rank.title }}
                      </span>
                    </td>
                    <td>{{ rank.members ? rank.members.length : '0' }}</td>
                    <td>{{ getMember(rank.id, 1)?.name }}</td>
                    <td>{{ getMember(rank.id, 2)?.name }}</td>
                    <td>{{ getMember(rank.id, 3)?.name }}</td>
                  </tr>
                </tbody>
              </v-table>
            </v-container>
          </v-main>
        </v-layout>
      </v-card>
    </v-container>

    <!-- Create Rank Dialog -->
    <div class="pa-4 text-center">
      <v-dialog v-model="createRankDialog" max-width="500">
        <v-card prepend-icon="mdi-database-plus-outline" title="Create Rank">
          <v-form method="post" action="/api/rank/create">
            <v-card-text>
              <v-row dense>
                <v-col cols="12">
                  <v-text-field name="title" label="Title" required />
                </v-col>
              </v-row>
            </v-card-text>

            <v-divider></v-divider>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn type="button" text="Close" variant="plain" @click="createRankDialog = false" />
              <v-btn type="submit" color="success" text="Save" variant="tonal" />
            </v-card-actions>
          </v-form>
        </v-card>
      </v-dialog>
    </div>

    <!-- Rank Members Dialog -->
    <v-dialog v-model="rankMembersDialog" max-width="500">
      <v-card>
        <v-card-title class="d-flex justify-space-between align-center">
          <span class="text-h6">{{ rank?.title }}</span>

          <v-btn icon variant="text">
            <v-icon>mdi-database-remove-outline</v-icon>
          </v-btn>
        </v-card-title>

        <v-container>
          <v-form ref="addMemberForm" @submit.prevent="submitAddOrRemoveMember">
            <v-row>
              <v-col cols="12">
                <v-text-field filled clearable type="text"
                  :append-icon="memberExists ? 'mdi-minus-thick' : 'mdi-plus-thick'" clear-icon="mdi-close-circle"
                  label="Name" v-model="memberName" @click:append="submitAddOrRemoveMember"
                  @input="onMemberNameInput" />
              </v-col>
            </v-row>
          </v-form>
        </v-container>

        <div class="flex-container">
          <div class="column">
            <div cols="6">
              <h4>Unranked</h4>
              <v-table>
                <draggable
                  class="list-group"
                  :component-data="{ tag: 'ul', name: 'flip-list', type: 'transition' }"
                  :list="unrankedMembers"
                  group="people"
                  @change="updateRanking"
                  itemKey="name"
                >
                  <template #item="{ element }">
                    <tr class="table-row-divider">
                      <td>{{ element.name }}</td>
                    </tr>
                  </template>
                </draggable>
              </v-table>
            </div>
          </div>

          <div class="column">
            <div cols="6">
              <h4>Ranked</h4>
              <v-table>
                <draggable class="list-group" :list="rankedMembers" group="people" @change="updateRanking"
                  itemKey="name">
                  <template #item="{ element, index }">
                    <tr class="table-row-divider">
                      <td>{{ index + 1 }} - {{ element.name }}</td>
                    </tr>
                  </template>
                </draggable>
              </v-table>
            </div>
          </div>
        </div>
      </v-card>
    </v-dialog>
  </div>
</template>

<script setup lang="ts">
import draggable from "vuedraggable"
import { onMounted, ref } from 'vue'
import { VBtn, VDialog, VCard, VCardTitle, VCardText, VCardActions, VSpacer } from 'vuetify/components'
import { getCookie } from "./token";

const createRankDialog = ref<boolean>(false)
const rankMembersDialog = ref<boolean>(false)
const token = ref<string | null>(null)
const rank = ref<Rank>()
const ranks = ref<Rank[]>([])
const members = ref<Member[]>([])
const unrankedMembers = ref<Member[]>([])
const rankedMembers = ref<Member[]>([])
const memberName = ref<string>("")
const memberExists = ref<boolean>(false)

onMounted(() => {
  checkToken()
  fetchRanks()
})

const checkToken = () => {
  token.value = getCookie("auth_token")
}

interface Rank {
  id: string;
  title: string;
  members: Member[];
}

interface Member {
  order: number;
  name: string;
}

const onMemberNameInput = (event: Event) => {
  ranks.value.some(rank =>
    rank.members.some(member => {
      if (member.name === memberName.value) {
        memberExists.value = true
      }
    })
  )
}

function getMember(rankId: string, orderNo: number): Member | undefined {
  const rank = ranks.value.find(r => r.id === rankId)
  if (rank) {
    return rank.members?.find(member => member.order === orderNo)
  }
  return undefined
}

function handleRanksRowClick(rankId: string): void { }

function handleShowRankMembersDialog(rankId: string) {
  rank.value = ranks.value.find(r => r.id === rankId)
  fetchRankMembers(rankId)
  rankMembersDialog.value = true
}

const addMemberForm = ref({
  rank_id: '',
  name: ''
})

const updateRanking = async (evt: any) => {
  console.log(evt)

  if (!rankedMembers.value) return

  const submission = {
    rank_id: rank.value!!.id,
    unranked_members: unrankedMembers.value,
    ranked_members: rankedMembers.value,
    event: evt,
  }

  console.log(submission)
  console.log(addMemberForm.value)

  if (addMemberForm.value) {
    try {
      const response = await fetch(`api/rank/sort`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(submission),
      })

      // Handle the response
      await response.json()

      fetchRank()
      fetchRanks()
    } catch (error) {
      console.error('Error submitting form:', error);
    }
  }
}

const submitAddOrRemoveMember = async () => {
  if (!rank.value?.id) { return }

  const submission = {
    rank_id: rank.value?.id,
    name: memberName.value,
  }

  if (addMemberForm.value) {
    const url = memberExists.value ? `/remove` : `/add`
    try {
      await fetch(`api/rank` + url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(submission),
      })

      // Reset the form after submission
      memberName.value = ''
      memberExists.value = false

      // Fetch the latest members and update the table
      fetchMembers()
    } catch (error) {
      console.error('Error submitting form:', error);
    }
  }
}

const fetchRank = async () => {
  try {
    const url = `/api/rank/query`
    const response = await fetch(`${url}?id=${rank.value?.id}`)
    if (!response.ok) {
      throw new Error('Bad networking...')
    }
    const data = await response.json()
    rank.value = data.rank as Rank
  } catch (err: any) {
    console.error(err)
  }
}

const fetchRanks = async () => {
  try {
    const url = `/api/rank`
    const response = await fetch(`${url}`)
    if (!response.ok) {
      throw new Error('Bad networking...')
    }
    const data = await response.json()
    ranks.value = data.ranks as Rank[]
  } catch (err: any) {
    console.error(err)
  }
}

const fetchRankMembers = async (rankId: string) => {
  try {
    const url = `/api/rank/members?rank_id=` + rankId
    const response = await fetch(`${url}`)
    if (!response.ok) {
      throw new Error('Bad networking...')
    }
    const data = await response.json()
    members.value = data.all_members as Member[]

    if (data.unranked_members) {
      unrankedMembers.value = data.unranked_members as Member[]
    } else {
      unrankedMembers.value = []
    }

    if (data.ranked_members) {
      rankedMembers.value = (data.ranked_members as Member[]).slice().sort((a: Member, b: Member) => a.order - b.order)
    } else {
      rankedMembers.value = []
    }
  } catch (err: any) {
    console.error(err)
  }
}

const fetchMembers = async () => {
  try {
    const url = `/api/rank/members?rank_id=` + rank.value?.id
    const response = await fetch(`${url}`)
    if (!response.ok) {
      throw new Error('Bad networking...')
    }
    const data = await response.json()
    members.value = data.all_members as Member[]

    if (data.unranked_members.length > 0) {
      unrankedMembers.value = data.unranked_members as Member[]
    } else {
      unrankedMembers.value = []
    }

    if (data.ranked_members.length > 0) {
      const ranking = data.ranked_members as Member[]
      rankedMembers.value = ranking.slice().sort((a: Member, b: Member) => a.order - b.order)
    } else {
      rankedMembers.value = []
    }
  } catch (err: any) {
    console.error(err)
  }
}
</script>

<style scoped>
.flex-container {
  display: flex;
  justify-content: center;
  align-items: flex-start;
  gap: 20px;
}

.column {
  flex: 1;
  padding: 10px;
}

.button {
  margin-top: 35px;
}

.flip-list-move {
  transition: transform 0.5s;
}

.no-move {
  transition: transform 0s;
}

.ghost {
  opacity: 0.5;
  background: #c8ebfb;
}

.list-group {
  min-height: 20px;
}

.list-group-item {
  cursor: move;
}

.list-group-item i {
  cursor: pointer;
}
</style>