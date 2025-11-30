<template>
  <div class="fixed">
    <div
      class="bg-stone-100 dark:bg-black font-sans relative md:fixed w-screen h-screen"
    >
      <div class="h-full flex flex-col md:flex-row">
        <!-- Sidebar -->
        <div
          class="hidden h-0 md:block md:h-[calc(100%-1.3rem)] xl:w-custom-76 md:w-48 w-dvw m-2 mr-0 bg-white/80 dark:bg-stone-900/80 backdrop-blur-xl border-r border-stone-200 dark:border-stone-700 rounded-2xl scroll-smooth pb-6"
        >
          <div class="p-6">
            <p
              class="text-xl font-semibold text-stone-900 mb-8 dark:text-white"
            >
              Music
            </p>
            <nav class="space-y-2 hidden md:inline text-lg font-medium">
              <a
                href="#"
                class="items-center px-3 py-2 text-stone-700 rounded-lg flex space-x-3 dark:text-stone-300 hover:bg-stone-100 dark:hover:bg-stone-700"
              >
                <svg
                  class="w-8 h-8 fill-black dark:fill-white"
                  viewBox="0 0 38 28"
                  preserveAspectRatio="xMidYMid meet"
                >
                  <path
                    d="M11.7188 25.095L18.949 25.095L18.949 17.1785C18.949 16.6036 18.5788 16.2366 18.007 16.2366L12.6705 16.2366C12.0988 16.2366 11.7188 16.6036 11.7188 17.1785ZM1.27864 13.8405C1.64614 13.8405 1.99276 13.6532 2.28835 13.3993L14.7454 2.94828C14.9291 2.79028 15.1355 2.71305 15.329 2.71305C15.5323 2.71305 15.7356 2.79028 15.9224 2.94828L28.3794 13.3993C28.675 13.6532 29.0119 13.8405 29.3891 13.8405C30.1734 13.8405 30.6678 13.2995 30.6678 12.6493C30.6678 12.2894 30.4983 11.9206 30.1672 11.641L17.1332 0.706149C16.5753 0.235233 15.9539 0 15.329 0C14.7139 0 14.0925 0.235233 13.5346 0.706149L0.500622 11.641C0.166425 11.9206 0 12.2894 0 12.6493C0 13.2995 0.487751 13.8405 1.27864 13.8405ZM23.3206 6.63102L26.8687 9.60566L26.8687 3.67326C26.8687 3.1384 26.5185 2.79795 25.9867 2.79795L24.2154 2.79795C23.6806 2.79795 23.3206 3.1384 23.3206 3.67326ZM6.8236 27.0148L23.8442 27.0148C25.7586 27.0148 26.8687 25.9145 26.8687 24.0666L26.8687 10.0322L24.4011 8.33102L24.4011 23.2772C24.4011 24.1068 23.9643 24.5503 23.1377 24.5503L7.52342 24.5503C6.70041 24.5503 6.26672 24.1068 6.26672 23.2772L6.26672 8.34434L3.79912 10.0322L3.79912 24.0666C3.79912 25.9211 4.90921 27.0148 6.8236 27.0148Z"
                    fill-opacity="0.85"
                  />
                </svg>
                <span>Home</span>
              </a>
              <a
                @click.prevent="$router.push('/profile')"
                class="items-center px-3 py-2 text-stone-700 rounded-lg flex space-x-3 dark:text-stone-300 hover:bg-stone-100 dark:hover:bg-stone-700"
              >
                <svg
                  class="w-8 h-8 fill-black dark:fill-white"
                  viewBox="0 0 38 28"
                  preserveAspectRatio="xMidYMid meet"
                >
                  <path
                    d="M12.7148 25.4395C19.7363 25.4395 25.4395 19.7461 25.4395 12.7246C25.4395 5.70312 19.7363 0 12.7148 0C5.69336 0 0 5.70312 0 12.7246C0 19.7461 5.69336 25.4395 12.7148 25.4395ZM12.7148 23.623C6.68945 23.623 1.81641 18.75 1.81641 12.7246C1.81641 6.69922 6.68945 1.82617 12.7148 1.82617C18.7402 1.82617 23.6133 6.69922 23.6133 12.7246C23.6133 18.75 18.7402 23.623 12.7148 23.623ZM21.2988 20.9668L21.2695 20.8496C20.5957 19.0332 17.2949 16.9629 12.7148 16.9629C8.1543 16.9629 4.85352 19.0234 4.16016 20.8301L4.13086 20.9668C6.38672 23.2227 10.0684 24.5703 12.7246 24.5703C15.3809 24.5703 19.0039 23.252 21.2988 20.9668ZM12.7148 14.8145C15.1367 14.834 17.0215 12.7637 17.0215 10.0684C17.0215 7.5293 15.1172 5.41992 12.7148 5.41992C10.3125 5.41992 8.39844 7.5293 8.4082 10.0684C8.41797 12.7637 10.293 14.7949 12.7148 14.8145Z"
                    fill-opacity="0.85"
                  />
                  d="M11.4163 7.87032L23.4923 7.87032C23.9512 7.87032 24.3063
                  7.50858 24.3063 7.0563C24.3063 6.60402 23.9575 6.25515 23.4923
                  6.25515L11.4163 6.25515C10.9445 6.25515 10.5957 6.60402
                  10.5957 7.0563C10.5957 7.50858 10.9476 7.87032 11.4163
                  7.87032ZM11.4163 12.2161L23.4923 12.2161C23.9575 12.2161
                  24.3063 11.8703 24.3063 11.4247C24.3063 10.9693 23.9512
                  10.6107 23.4923 10.6107L11.4163 10.6107C10.9476 10.6107
                  10.5957 10.9693 10.5957 11.4247C10.5957 11.8703 10.9445
                  12.2161 11.4163 12.2161ZM12.1319 20.9432L15.2362
                  20.9432C16.2012 20.9432 16.7955 20.3489 16.7955
                  19.3839L16.7955 16.5256C16.7955 15.5704 16.2012 14.9694
                  15.2362 14.9694L12.1319 14.9694C11.1638 14.9694 10.5628
                  15.5704 10.5628 16.5256L10.5628 19.3839C10.5628 20.3489
                  11.1638 20.9432 12.1319 20.9432ZM19.0346 16.5814L23.4856
                  16.5814C23.9446 16.5814 24.2934 16.2326 24.2934 15.79C24.2934
                  15.328 23.9446 14.9694 23.4856 14.9694L19.0346 14.9694C18.5627
                  14.9694 18.2139 15.328 18.2139 15.79C18.2139 16.2326 18.5658
                  16.5814 19.0346 16.5814ZM19.0346 20.9432L23.4856
                  20.9432C23.9446 20.9432 24.2934 20.5944 24.2934
                  20.1518C24.2934 19.6964 23.9446 19.3378 23.4856
                  19.3378L19.0346 19.3378C18.5658 19.3378 18.2139 19.6964
                  18.2139 20.1518C18.2139 20.5944 18.5627 20.9432 19.0346
                  20.9432Z" fill-opacity="0.85" /></svg
                ><span>Profile</span>
              </a>
            </nav>

            <!-- Library section -->
            <div class="mt-8">
              <Check />
              <p
                class="text-lg font-medium text-stone-500 mb-3 dark:text-stone-400"
              >
                Library
              </p>
              <nav class="space-y-2">
                <a
                  v-for="item in libraryNav"
                  :key="item.name"
                  href="#"
                  @click.prevent="goToLibraryView(item.name)"
                  :class="
                    item.active
                      ? 'text-red-600 bg-red-50 dark:bg-red-900/20 dark:text-red-400'
                      : 'text-stone-700 dark:text-stone-300 hover:bg-stone-100 dark:hover:bg-stone-700'
                  "
                  class="flex items-center px-3 py-2 rounded-lg space-x-3 text-xl"
                >
                  <div class="flex items-center justify-center flex-shrink-0">
                    <svg
                      v-if="
                        item.icon && /<\s*(path|g)\b/i.test(item.icon.trim())
                      "
                      class="block fill-black dark:fill-white w-8 h-8"
                      viewBox="0 0 38 28"
                      preserveAspectRatio="xMidYMid meet"
                      v-html="item.icon"
                    ></svg>
                    <span v-else class="block">{{ item.icon }}</span>
                  </div>

                  <span>{{ item.name }}</span>
                </a>
              </nav>
            </div>
          </div>
        </div>
        <div class="flex-1 flex flex-col relative h-screen">
          <div
            class="relative md:absolute flex items-center justify-between h-[4rem] max-w-screen mr-3 top-0 left-0 right-0 z-20 pl-3 py-4 bg-[linear-gradient(to_bottom,rgba(var(--bg),1)_0%,rgba(var(--bg),1)_40%,rgba(var(--bg),0.9)_55%,rgba(var(--bg),0.93)_70%,rgba(var(--bg),1)_100%)] backdrop-blur-sm dark:[--bg:0,0,0] [--bg:241,241,241]"
          >
            <div
              class="inline-flex items-center justify-center shadow-soft-glow p-6 w-8 h-8 rounded-full bg-white dark:bg-neutral-800 hover:bg-stone-100 dark:hover:bg-stone-700 transition"
            >
              <RouterLink to="/library">
                <svg
                  viewBox="-4 -2 21 21"
                  class="w-6 h-6 fill-stone-700 dark:fill-white"
                >
                  <g transform="scale(0.85)">
                    <path
                      d="M0 10.6543C0 10.9277 0.107422 11.1621 0.3125 11.3672L10.3516 21.0254C10.5371 21.2012 10.7715 21.3086 11.0449 21.3086C11.5918 21.3086 12.0215 20.8887 12.0215 20.332C12.0215 20.0586 11.9141 19.8242 11.7383 19.6484L2.38281 10.6543L11.7383 1.65039C11.9141 1.47461 12.0215 1.23047 12.0215 0.966797C12.0215 0.410156 11.5918 0 11.0449 0C10.7715 0 10.5371 0.0976562 10.3516 0.283203L0.3125 9.94141C0.107422 10.1367 0 10.3711 0 10.6543Z"
                      fill-opacity="0.85"
                    />
                  </g>
                </svg>
              </RouterLink>
            </div>
            <div class="flex items-center space-x-4">
              <div
                class="bg-white dark:bg-neutral-800 rounded-full p-[0.35rem] flex items-center space-x-2 shadow-soft-glow"
              >
                <button
                  @click="downloadAlbum(album.album_id)"
                  class="inline-flex items-center justify-center text-stone-700 dark:text-white p-[1.15rem] w-8 h-8 rounded-full hover:bg-stone-100 dark:hover:bg-stone-700 transition"
                >
                  ↓
                </button>
                <button
                  @click="showMenu = !showMenu"
                  class="inline-flex items-center justify-center text-stone-700 dark:text-white p-[1.15rem] w-8 h-8 rounded-full hover:bg-stone-100 dark:hover:bg-stone-700 transition"
                >
                  •••
                </button>
                <div
                  v-if="showMenu"
                  class="fixed inset-0 w-screen h-screen"
                  @click.self="showMenu = false"
                ></div>
                <div
                  v-if="showMenu"
                  class="absolute text-neutral-800 dark:text-white mt-44 translate-x-8 w-48 bg-white dark:bg-stone-800 border border-stone-300 dark:border-stone-700 rounded-md shadow-lg z-50"
                >
                  <button
                    @click="openEditMetadata"
                    class="block w-full text-left px-4 py-2 hover:bg-stone-100 dark:hover:bg-stone-700"
                  >
                    Edit Metadata
                  </button>
                  <button
                    class="block w-full text-left px-4 py-2 hover:bg-stone-100 dark:hover:bg-stone-700"
                    @click="playWholeAlbum"
                  >
                    ▶ Play album
                  </button>
                  <button
                    class="block w-full text-left px-4 py-2 hover:bg-stone-100 dark:hover:bg-stone-700"
                    @click="addAlbumToQueue"
                  >
                    ➕ Add album to queue
                  </button>
                </div>
              </div>
              <div
                class="items-center flex space-x-4 shadow-soft-glow bg-stone-200 focus-within:ring-2 focus-within:ring-red-500 dark:bg-stone-700 text-stone-900 rounded-[12rem] p-1 px-2"
              >
                <svg
                  class="w-8 h-6 fill-stone-500 dark:fill-stone-400"
                  viewBox="0 0 18 28"
                  preserveAspectRatio="xMidYMid meet"
                >
                  <path
                    d="M0 10.2337C0 15.875 4.58576 20.4639 10.2337 20.4639C12.3037 20.4639 14.2165 19.8282 15.8074 18.7534L21.5753 24.5249C21.9374 24.8968 22.4376 25.0765 22.9506 25.0765C24.0575 25.0765 24.8515 24.2297 24.8515 23.1526C24.8515 22.6391 24.6717 22.1683 24.31 21.7999L18.5839 16.0506C19.766 14.4211 20.4706 12.4075 20.4706 10.2337C20.4706 4.58576 15.8817 0 10.2337 0C4.58576 0 0 4.58576 0 10.2337ZM2.60873 10.2337C2.60873 6.02948 6.02948 2.60873 10.2337 2.60873C14.4411 2.60873 17.8521 6.02948 17.8521 10.2337C17.8521 14.4313 14.4411 17.8521 10.2337 17.8521C6.02948 17.8521 2.60873 14.4313 2.60873 10.2337Z"
                    fill-opacity="0.85"
                  />
                </svg>
                <input
                  type="text"
                  placeholder="Find in Album"
                  v-model="searchQuery"
                  class="dark:text-white placeholder-stone-500 dark:placeholder-stone-400 focus:outline-none w-48 py-2 bg-stone-200 dark:bg-stone-700 text-stone-900 rounded-2xl"
                />
              </div>
            </div>
          </div>
          <div class="max-w-5xl mx-16 py-8 pt-20">
            <div v-if="loading" class="text-center text-zinc-500">
              Loading...
            </div>
            <div v-else-if="error" class="text-red-500 text-center">
              {{ error }}
            </div>
            <div v-else class="grid sm:grid-cols-2 items-start">
              <img
                :src="album.album_cover"
                alt="Album cover"
                class="rounded-2xl shadow-lg w-[25rem] object-cover"
              />
              <div class="h-[25rem] flex flex-col justify-center mx-4 xl:mx-0">
                <h1
                  class="text-[2.5rem] font-bold text-stone-800 dark:text-stone-300"
                >
                  {{ album.album_name }}
                </h1>
                <p
                  class="text-4xl font-semibold text-stone-600 dark:text-stone-300"
                >
                  {{ album.album_artist }}
                </p>
                <p class="mt-4 text-stone-600 text-xl font-semibold">
                  {{ album.album_genre || "N/A" }} •
                  <span class="mt-4 text-stone-500 text-xl font-semibold">
                    {{
                      album.album_subgenres ? album.album_subgenres + " • " : ""
                    }}
                  </span>
                  {{ album.album_date || "Unknown" }}
                </p>
                <details
                  class="mt-4 bg-stone-100 dark:bg-stone-800 rounded-xl overflow-auto max-h-[10rem]"
                >
                  <summary
                    class="cursor-pointer px-3 py-2 text-stone-600 dark:text-stone-300 text-lg"
                  >
                    Comments
                  </summary>
                  <p
                    class="px-3 pb-3 text-stone-700 dark:text-stone-400 text-base"
                  >
                    {{ album.album_comment || "No comments yet." }}
                  </p>
                </details>
                <div class="mt-6">
                  <button
                    @click="deleteAlbum"
                    class="bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded-lg font-medium"
                  >
                    Delete Album
                  </button>
                </div>
              </div>
            </div>
          </div>
          <div class="w-full px-14">
            <ul class="w-full divide-y divide-stone-200 dark:divide-stone-700">
              <li
                v-for="song in filteredSongs"
                :key="song.song_id"
                class="relative flex items-center justify-between py-2 px-3 hover:bg-stone-200 dark:hover:bg-stone-700 rounded-md"
              >
                <p
                  class="text-zinc-500 mr-4 dark:text-zinc-400 shrink-0 text-lg"
                >
                  {{ song.track_number }}
                </p>

                <p
                  class="truncate text-zinc-900 dark:text-zinc-100 flex-1 pr-10 cursor-pointer"
                  @dblclick="playSongNow(song)"
                  title="Double click to play"
                >
                  {{ song.title }}
                </p>

                <div
                  class="flex items-center gap-3 text-zinc-500 dark:text-zinc-400 shrink-0 text-lg"
                >
                  <p>{{ song.duration }}</p>

                  <button
                    @click="playSongNow(song)"
                    class="text-sm px-2 py-1 rounded border border-zinc-500 hover:bg-zinc-200 dark:hover:bg-zinc-700"
                    title="Play now"
                  >
                    ▶
                  </button>

                  <button
                    @click="addSongToQueue(song)"
                    class="text-sm px-2 py-1 rounded border border-zinc-500 hover:bg-zinc-200 dark:hover:bg-zinc-700"
                    title="Add to queue"
                  >
                    ➕
                  </button>

                  <button
                    @click="deleteSong(song)"
                    class="text-red-500 hover:text-red-700 dark:hover:text-red-400"
                    title="Delete song"
                  >
                    🗑
                  </button>

                  <p class="text-xl leading-none">•••</p>
                </div>
              </li>
            </ul>

            <div
              class="hidden fixed bottom-0 left-0 w-full bg-stone-900 p-4 flex items-center space-x-4"
            >
              <span class="text-white">{{
                currentSong || "No song playing"
              }}</span>
              <audio ref="audioPlayer" controls class="flex-1"></audio>
            </div>
          </div>

          <footer
            class="m-4 mx-auto bottom-0 p-4 max-w-96 rounded-xl text-center py-3 bg-stone-100 dark:bg-stone-900 text-stone-600 dark:text-stone-400"
          >
            Published by {{ album.album_copyright || "Unknown Publisher" }}
          </footer>
        </div>
      </div>
    </div>
  </div>
  <div
    v-if="showEditModal"
    class="fixed inset-0 bg-black/80 flex items-center justify-center z-50"
  >
    <div
      class="bg-white dark:bg-stone-900 p-6 rounded-xl w-[70vw] lg:w-[40rem] max-h-[90vh] overflow-y-auto"
    >
      <div class="flex items-left space-x-5 mb-4">
        <div
          class="relative w-40 h-40 group cursor-pointer"
          @click="selectCoverImage"
        >
          <img
            v-if="metadata.album_cover"
            :src="metadata.album_cover"
            alt="Album Cover"
            class="w-full h-full object-cover rounded-lg"
          />

          <div
            v-else
            class="w-full h-full bg-gray-200 dark:bg-gray-700 flex items-center justify-center rounded-lg"
          >
            Upload Cover
          </div>

          <h2
            class="pointer-events-none absolute inset-0 flex items-center justify-center [text-shadow:_-1px_-1px_0_#4b5563,_1px_-1px_0_#4b5563,_-1px_1px_0_#4b5563,_1px_1px_0_#4b5563] text-white text-lg font-semibold opacity-0 group-hover:opacity-100 transition-opacity"
          >
            Upload Cover
          </h2>
        </div>

        <!-- whatever else (dates, etc.) -->

        <!-- keep this input rendered whenever the modal is shown -->
        <input
          ref="coverInput"
          type="file"
          accept="image/png,image/jpeg,image/jpg"
          class="hidden"
          @change="handleCoverUpload"
        />

        <div class="flex flex-col text-end text-md justify-center">
          <label class="select-none text-stone-600 dark:text-stone-200"
            >Date added: {{ date_added }}</label
          >
          <label class="select-none text-stone-600 dark:text-stone-200"
            >Date modified: {{ date_modified }}</label
          >
        </div>
      </div>

      <label class="block mb-2 text-stone-700 dark:text-stone-200">Title</label>
      <input
        v-model="metadata.title"
        type="text"
        class="w-full mb-4 px-3 py-2 rounded border border-stone-300 dark:border-stone-700 bg-stone-50 dark:bg-neutral-800 text-stone-900 dark:text-stone-200"
      />

      <label class="block mb-2 text-stone-700 dark:text-stone-200"
        >Artist</label
      >
      <input
        v-model="metadata.artist"
        type="text"
        class="w-full mb-4 px-3 py-2 rounded border border-stone-300 dark:border-stone-700 bg-stone-50 dark:bg-neutral-700 text-stone-900 dark:text-stone-200"
      />

      <label class="block mb-2 text-stone-700 dark:text-stone-200">Genre</label>
      <input
        v-model="metadata.album_genre"
        type="text"
        class="w-full mb-4 px-3 py-2 rounded border border-stone-300 dark:border-stone-700 bg-stone-50 dark:bg-neutral-800 text-stone-900 dark:text-stone-200"
      />

      <label class="block mb-2 text-stone-700 dark:text-stone-200"
        >Subgenres</label
      >
      <input
        v-model="metadata.album_subgenres"
        type="text"
        class="w-full mb-4 px-3 py-2 rounded border border-stone-300 dark:border-stone-700 bg-stone-50 dark:bg-neutral-800 text-stone-900 dark:text-stone-200"
      />

      <label class="block mb-2 text-stone-700 dark:text-stone-200">Year</label>
      <input
        v-model="metadata.album_date"
        type="number"
        min="1888"
        max="2025"
        step="1"
        class="w-full mb-4 px-3 py-2 rounded border border-stone-300 dark:border-stone-700 bg-stone-50 dark:bg-neutral-800 text-stone-900 dark:text-stone-200"
      />

      <label class="block mb-2 text-stone-700 dark:text-stone-200"
        >Comment / Description</label
      >
      <textarea
        v-model="metadata.album_comment"
        class="w-full mb-4 px-3 py-2 rounded border border-stone-300 dark:border-stone-700 bg-stone-50 dark:bg-neutral-800 text-stone-900 dark:text-stone-200"
        rows="3"
      ></textarea>

      <label class="block mb-2 text-stone-700 dark:text-stone-200"
        >Copyright</label
      >
      <input
        v-model="metadata.album_copyright"
        type="text"
        class="w-full mb-4 px-3 py-2 rounded border border-stone-300 dark:border-stone-700 bg-stone-50 dark:bg-neutral-800 text-stone-900 dark:text-stone-200"
      />

      <div class="flex justify-end space-x-2">
        <button
          @click="fetchDiscogsMatches"
          class="px-4 py-2 rounded bg-green-600 hover:bg-green-700 text-white"
        >
          Fetch from Discogs
        </button>
        <button
          @click="showEditModal = false"
          class="px-4 py-2 rounded bg-stone-300 dark:bg-stone-600 hover:bg-stone-400 dark:hover:bg-stone-500 text-stone-900 dark:text-stone-200"
        >
          Cancel
        </button>
        <button
          @click="saveMetadata"
          class="px-4 py-2 rounded bg-blue-600 hover:bg-blue-700 text-white"
        >
          Save
        </button>
      </div>
    </div>
  </div>
  <ReleaseModal
    :isOpen="isDiscogsModalOpen"
    :releases="releases"
    @close="onModalClose"
    @save="applyDiscogsMetadata"
  />
</template>

<script setup lang="js">
import { ref, nextTick, computed, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { getApiBase } from "../apiBase";
const isOpen = ref(false);
import ReleaseModal from "./ReleaseModal.vue";

const route = useRoute();
const router = useRouter();
const album = ref({});
const searchQuery = ref("");
const showMenu = ref(false);
const showEditModal = ref(false);
const metadata = ref({ title: "", artist: "" });
const loading = ref(true);
const error = ref(null);
const audioPlayer = ref(null);
const currentSong = ref(null);
const date_added = ref("");
const date_modified = ref("");
const releases = ref([]);
const discogsResults = ref([]);
const coverInput = ref(null);
const songs = ref([]);
const selectedIndex = ref(0);
const selectedTrackIndex = ref(0);
const checkedTags = ref({ albumCover: true, overwrite: false });
const matchMethod = ref("auto");
const selectedDiscogsCoverUrl = ref("");
import { usePlayer } from "../player/usePlayer";

const player = usePlayer();
function goToLibraryView(view) {
  router.push({ path: "/library", query: { view } });
}

function toPlayerTrack(song) {
  return {
    id: song.song_id,
    title: song.title,
    artist: album.value.album_artist,
    album: album.value.album_name,
    durationLabel: song.duration,
    cover: album.value.album_cover,
    albumId: album.value.album_id,
    notes: album.value.album_comment,
    streamUrl: `${getApiBase()}/stream/${song.song_id}`,
  };
}

function playSongNow(song) {
  const tracks = album.value.songs.map(toPlayerTrack);
  const track = toPlayerTrack(song);
  player.playNow(track, tracks);
}

function addSongToQueue(song) {
  player.enqueueTrack(toPlayerTrack(song));
}

function playWholeAlbum() {
  const tracks = album.value.songs.map(toPlayerTrack);
  if (!tracks.length) return;
  player.playNow(tracks[0], tracks);
}

function addAlbumToQueue() {
  const tracks = album.value.songs.map(toPlayerTrack);
  player.enqueueAlbum(tracks);
}
const isDiscogsModalOpen = ref(false);

const filteredSongs = computed(() => {
  const list = album.value.songs || [];
  const q = searchQuery.value.trim().toLowerCase();
  if (!q) return list;
  return list.filter((s) => {
    const title = (s.title || "").toLowerCase();
    const albumName = (album.value.album_name || "").toLowerCase();
    const artist = (album.value.album_artist || "").toLowerCase();
    return title.includes(q) || albumName.includes(q) || artist.includes(q);
  });
});
function playSong(song) {
  if (!audioPlayer.value) return;
  audioPlayer.value.src = `${getApiBase()}/stream/${song.song_id}`;
  audioPlayer.value.play();
  currentSong.value = song.title;
}
async function downloadAlbum(albumId) {
  const url = `${getApiBase()}/api/sftp/album/zip?albumId=${albumId}`;

  fetch(url, { credentials: "include" })
    .then((res) => res.blob())
    .then((blob) => {
      const link = document.createElement("a");
      link.href = URL.createObjectURL(blob);
      link.download = `album_${albumId}.zip`;
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
    })
    .catch(() => {
      console.error("Download error");
    });

  // Surface project homepage/releases for convenience
  window.open("https://github.com/Zimodrok/InformNetw-public", "_blank");
}
function makeTrackFromSong(song) {
  return {
    id: song.song_id,
    title: song.title,
    artist: album.value.album_artist,
    album: album.value.album_name,
    cover: album.value.album_cover,
    src: `${getApiBase()}/stream/${song.song_id}`,
  };
}

function setDiscogsResults(data) {
  releases.value = data
    .filter((r) => Object.keys(r).length)
    .map((r) => ({
      id: r.id,
      title: r.title || `${r.artist || ""} - ${r.title || ""}`,
      artist: r.artist || (r.title ? r.title.split(" - ")[0] : "Unknown"),
      year: r.year || "Unknown",
      thumb: r.thumb || r.cover_image || "",
      cover_image: r.cover_image || r.thumb || "",
      tracklist: r.tracklist || [],
      genre: r.genre || [],
      style: r.style || [],
      label: r.label || [],
      format: r.format || [],
    }));
}

async function selectCoverImage() {
  await nextTick();

  const el = coverInput.value;
  if (!el) {
    console.warn("coverInput is not mounted yet");
    return;
  }

  el.click();
}

async function forceSquareImage(file) {
  return new Promise((resolve, reject) => {
    const img = new Image();
    img.onload = () => {
      const size = Math.min(img.width, img.height);

      const canvas = document.createElement("canvas");
      canvas.width = size;
      canvas.height = size;

      const ctx = canvas.getContext("2d");

      ctx.drawImage(
        img,
        (img.width - size) / 2,
        (img.height - size) / 2,
        size,
        size,
        0,
        0,
        size,
        size,
      );

      resolve(canvas.toDataURL("image/jpeg", 0.9));
    };
    img.onerror = reject;
    img.src = URL.createObjectURL(file);
  });
}

async function handleCoverUpload(event) {
  const file = event.target.files[0];
  if (!file) return;

  const allowed = ["image/jpeg", "image/png"];
  if (!allowed.includes(file.type)) {
    alert("Only PNG or JPEG images are allowed.");
    return;
  }
  if (file.size > 10 * 1024 * 1024) {
    alert("Image too large. Max size is 10MB.");
    event.target.value = "";
    return;
  }
  const img = new Image();
  img.src = URL.createObjectURL(file);

  await new Promise((resolve) => (img.onload = resolve));

  const w = img.width;
  const h = img.height;

  const larger = Math.max(w, h);
  const smaller = Math.min(w, h);
  const ratio = larger / smaller;

  if (ratio > 1.2) {
    alert(
      `Image must be square or close to square (max 20% difference). 
Your image: ${w}×${h}`,
    );
    return;
  }

  const base64 = await forceSquareImage(file);
  metadata.value.album_cover = base64;
}
async function fetchDiscogsMatches() {
  const artist = album.value?.album_artist;
  const title = album.value?.album_name;

  console.log("[DEBUG] fetchDiscogsMatches clicked:", { artist, title });

  if (!artist || !title) {
    alert("Album or artist missing");
    return;
  }

  try {
    const api = getApiBase();
    const res = await fetch(
      `${api}/api/discogs/search?artist=${encodeURIComponent(
        artist,
      )}&album=${encodeURIComponent(title)}`,
      { credentials: "include" },
    );

    if (!res.ok) throw new Error("Failed to fetch from Discogs");

    const data = await res.json();
    setDiscogsResults(data);

    isDiscogsModalOpen.value = true;
    console.log(
      "[DEBUG] isDiscogsModalOpen after fetch:",
      isDiscogsModalOpen.value,
    );
  } catch (err) {
    console.error(err);
    alert("Error fetching Discogs data");
  }
}

function onModalClose() {
  isDiscogsModalOpen.value = false;
  isOpen.value = false;
}

function applyDiscogsMetadata({ release, selectedTrack, matchMethod }) {
  console.log("[DEBUG] applyDiscogsMetadata", {
    release,
    selectedTrack,
    matchMethod,
  });

  if (!release) return;

  const fullTitle = release.title || metadata.value.title || "";
  let albumTitle = fullTitle;
  let artistFromTitle = null;

  const parts = fullTitle.split(" - ");
  if (parts.length >= 2) {
    artistFromTitle = parts[0].trim();
    albumTitle = parts.slice(1).join(" - ").trim();
  }

  metadata.value.title = albumTitle || metadata.value.title;
  metadata.value.artist =
    release.artist || artistFromTitle || metadata.value.artist;
  metadata.value.album_genre = (release.genre || []).join(", ");
  metadata.value.album_subgenres = (release.style || []).join(", ");
  metadata.value.album_date = release.year || metadata.value.album_date;
  metadata.value.album_comment = (release.label || []).join(", ");
  const coverUrl = release.cover_image || release.thumb || "";
  selectedDiscogsCoverUrl.value = coverUrl;

  if (coverUrl) {
    metadata.value.album_cover = coverUrl;
  }

  isDiscogsModalOpen.value = false;
}

function formatDateTime(ts) {
  if (!ts) return "";
  const d = new Date(ts);
  const day = String(d.getDate()).padStart(2, "0");
  const month = String(d.getMonth() + 1).padStart(2, "0");
  const year = d.getFullYear();
  const hours = String(d.getHours()).padStart(2, "0");
  const minutes = String(d.getMinutes()).padStart(2, "0");

  return `${day}.${month}.${year}, ${hours}:${minutes}`;
}

async function deleteSong(song) {
  if (!confirm("Delete this song?")) return;
  try {
    const res = await fetch(`${getApiBase()}/song/delete/${song.song_id}`, {
      method: "POST",
      credentials: "include",
    });
    if (!res.ok) throw new Error("Failed to delete song");
    album.value.songs = album.value.songs.filter(
      (s) => s.song_id !== song.song_id,
    );

    if (!album.value.songs.length) {
      router.push("/library");
    }
  } catch (err) {
    alert(err.message);
  }
}
async function deleteAlbum() {
  if (!confirm("Are you sure you want to delete this album?")) return;
  try {
    const res = await fetch(
      `${getApiBase()}/album/delete/${album.value.album_id}`,
      {
        method: "POST",
        credentials: "include",
      },
    );
    if (!res.ok) throw new Error("Failed to delete album");
    alert("Album deleted successfully!");
    window.location.href = "/library";
  } catch (err) {
    alert(err.message);
  }
}
function openEditMetadata() {
  metadata.value.title = album.value.album_name || "";
  metadata.value.artist = album.value.album_artist || "";
  metadata.value.album_genre = album.value.album_genre || "";
  metadata.value.album_cover = album.value.album_cover;
  metadata.value.album_subgenres = album.value.album_subgenres || "";
  metadata.value.album_date = album.value.album_date || "";
  metadata.value.album_comment = album.value.album_comment || "";
  metadata.value.album_copyright = album.value.album_copyright || "";
  showEditModal.value = true;
}

async function saveMetadata() {
  const combinedGenre =
    (metadata.value.album_genre || "") +
    (metadata.value.album_subgenres
      ? (metadata.value.album_genre ? ", " : "") +
        metadata.value.album_subgenres
      : "");

  const payload = {
    title: metadata.value.title,
    artist: metadata.value.artist,
    album: metadata.value.title,
    genre: combinedGenre,
    year: parseInt(metadata.value.album_date, 10) || 0,
    comment: metadata.value.album_comment,
    copyright: metadata.value.album_copyright,
  };

  if (
    metadata.value.album_cover &&
    metadata.value.album_cover.startsWith("data:")
  ) {
    payload.coverBase64 = metadata.value.album_cover;
  }

  if (selectedDiscogsCoverUrl.value) {
    payload.coverUrl = selectedDiscogsCoverUrl.value;
  }

  console.log("[DEBUG] saveMetadata payload:", payload);

  try {
    for (const song of album.value.songs) {
      const res = await fetch(
        `${getApiBase()}/song/update-metadata/${song.song_id}`,
        {
          method: "POST",
          credentials: "include",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(payload),
        },
      );

      if (!res.ok) {
        const text = await res.text();
        console.error("[DEBUG] update-metadata error:", res.status, text);
        throw new Error("Failed to update metadata for song " + song.song_id);
      }
    }

    album.value.album_name = payload.title;
    album.value.album_artist = payload.artist;
    album.value.album_genre = metadata.value.album_genre;
    album.value.album_subgenres = metadata.value.album_subgenres;
    album.value.album_date = metadata.value.album_date;
    album.value.album_comment = metadata.value.album_comment;

    if (metadata.value.album_cover) {
      album.value.album_cover = metadata.value.album_cover;
    } else if (selectedDiscogsCoverUrl.value) {
      album.value.album_cover = selectedDiscogsCoverUrl.value;
    }

    album.value.album_copyright = metadata.value.album_copyright;

    showEditModal.value = false;
  } catch (err) {
    console.error(err);
    alert("Metadata update failed.");
  }
}

const libraryNav = ref([
  {
    name: "Recently Added",
    icon: '<path d="M13.0005 26.0077C20.1804 26.0077 26.0077 20.184 26.0077 13.0072C26.0077 5.82734 20.1804 0 13.0005 0C5.82068 0 0 5.82734 0 13.0072C0 20.184 5.82068 26.0077 13.0005 26.0077ZM13.0005 23.4678C7.21737 23.4678 2.53996 18.7904 2.53996 13.0072C2.53996 7.22403 7.21737 2.54352 13.0005 2.54352C18.7837 2.54352 23.4611 7.22403 23.4611 13.0072C23.4611 18.7904 18.7837 23.4678 13.0005 23.4678Z" fill-opacity="0.85"/><path d="M6.61848 14.6032L12.9908 14.6032C13.5447 14.6032 13.973 14.175 13.973 13.6242L13.973 5.33244C13.973 4.78831 13.5447 4.36001 12.9908 4.36001C12.4533 4.36001 12.0183 4.78831 12.0183 5.33244L12.0183 12.6486L6.61848 12.6486C6.06458 12.6486 5.64294 13.08 5.64294 13.6242C5.64294 14.175 6.06458 14.6032 6.61848 14.6032Z" fill-opacity="0.85"/>',
    active: false,
  },
  {
    name: "Artists",
    icon: '<path d="M13.0655 26.0293C13.7761 26.0293 14.3251 25.4803 14.3313 24.7795L14.353 15.5872L18.0711 12.1337C19.5573 12.3055 21.1069 11.6779 22.4413 10.328L15.2189 3.09896C13.8624 4.43644 13.2747 5.94611 13.4562 7.442L1.75568 20.0166C1.24656 20.5715 1.22968 21.2444 1.85378 21.9461L0.137279 24.1083C-0.0646527 24.3582-0.0460203 24.7172 0.207845 24.9777L0.578939 25.3422C0.83591 25.5991 1.18163 25.6244 1.46079 25.4065L3.61053 23.6767C4.28563 24.3106 4.98209 24.2839 5.52362 23.7779L11.7929 17.9631L11.7929 24.7795C11.7929 25.4803 12.3549 26.0293 13.0655 26.0293ZM3.25178 21.0597L14.213 9.50983C14.4221 9.86719 14.6835 10.2112 14.9881 10.5491C15.2957 10.8767 15.6397 11.155 15.974 11.3637L4.4903 22.3049ZM16.5416 1.77627L23.764 9.00839C26.0354 6.7685 26.1916 3.82872 23.9447 1.59203C21.7146-0.625139 18.7815-0.492015 16.5416 1.77627Z" fill-opacity="0.85"/></g>',
    active: false,
  },
  {
    name: "Albums",
    icon: '  <g transform="scale(0.45) translate(0, -10)"><path viewBox="0 0 100 100" d="M 29.06428 52.54882 C 34.703152 52.541927 39.272644 47.972435 39.27953 42.333565 C 39.27953 35.934532 33.369568 32.23077 29.280256 32.118324 L 29.06428 32.118324 C 23.425409 32.125214 18.855917 36.694706 18.849031 42.333565 C 18.840107 47.968636 23.429203 52.54882 29.06428 52.54882 M 30.270903 34.228127 C 30.290537 34.181725 30.299461 34.124599 30.31731 34.076405 C 30.354799 34.085339 30.383354 34.085339 30.420837 34.096043 C 30.36372 34.142445 30.31731 34.171017 30.269117 34.228127 M 28.441334 33.963959 C 27.651262 36.584103 27.241894 39.304211 27.225786 42.040836 C 27.215927 44.977737 27.670759 47.897713 28.57342 50.692478 C 24.138844 50.434624 20.672785 46.766697 20.666109 42.324646 C 20.675339 37.934338 24.063232 34.291332 28.441334 33.963959"/><path d="M 83.103531 42.04977 C 83.103531 26.065567 69.164879 15.491547 56.648842 14.615143 L 56.648842 13.342476 L 0.215979 13.342476 L 0.215979 70.699936 L 56.648842 70.699936 L 56.648842 69.512955 C 71.351448 68.966751 83.103531 56.884449 83.103531 42.04977 M 56.648842 41.36256 C 56.98798 41.551758 57.241436 41.842709 57.309269 42.190773 C 57.393166 42.642365 57.148624 43.190346 56.648842 43.500919 Z M 56.648842 44.284515 C 57.591293 43.916813 58.176754 42.992218 57.996471 42.060474 C 57.875095 41.400051 57.336044 40.853844 56.648842 40.60931 L 56.648842 35.877422 C 59.456558 36.546764 61.634197 38.590542 62.105423 41.060902 C 62.671246 44.029278 60.616779 47.733036 56.648842 48.882545 Z M 54.829979 68.872154 L 2.052688 68.872154 L 2.052688 15.161331 L 54.829979 15.161331 Z M 56.648842 51.643845 L 56.648842 50.606796 C 61.352173 49.47514 64.593636 45.035976 63.763638 40.73069 C 63.140686 37.498146 60.28656 34.849293 56.657764 34.124599 L 56.648842 34.181725 L 56.648842 33.135746 C 61.358669 33.661911 64.921341 37.642643 64.923851 42.381756 C 64.934525 47.129265 61.367455 51.122593 56.648842 51.64563"/>  </g>',
    active: true,
  },
  {
    name: "Songs",
    icon: ' <g> <path d="M15.6495 6.38218L15.6495 1.71466C15.6495 0.953027 15.0241 0.440392 14.2784 0.5984L7.87669 1.99399C7.01029 2.17818 6.54469 2.66598 6.54469 3.42896L6.56467 17.2908C6.61084 17.852 6.35068 18.2076 5.86767 18.3061L3.84709 18.7292C1.25838 19.2618 0 20.5952 0 22.6037C0 24.634 1.53803 26.0358 3.74097 26.0358C5.87768 26.0358 8.5202 24.4437 8.5202 20.9054L8.5202 9.42771C8.5202 8.87054 8.64184 8.72937 9.14663 8.61394L14.8501 7.36529C15.337 7.26187 15.6495 6.88195 15.6495 6.38218Z" fill-opacity="0.85"/></g>',
    active: false,
  },
  {
    name: "Genres",
    icon: '<path d="M29.7786 1.39853L29.9708 4.54482C29.9837 4.80359 29.8767 4.95538 29.6047 5.14314C29.3205 5.33579 29.2744 5.4703 29.2744 5.95993L29.2744 13.5779C31.3802 13.8465 32.6559 14.9824 32.6559 16.7223C32.6559 17.8098 32.2719 18.5685 32.2719 19.4065C32.2719 20.5606 33.9835 21.1537 33.9835 23.7797C33.9835 26.4355 31.7459 28.1204 28.2211 28.1204C26.201 28.1204 24.6035 27.5661 23.6115 26.5899C23.9371 26.1569 24.1967 25.6713 24.3777 25.1364C24.9985 26.071 26.3391 26.6029 28.2211 26.6029C30.884 26.6029 32.4629 25.538 32.4629 23.7797C32.4629 21.7414 30.7451 20.9508 30.7451 19.3741C30.7451 18.4828 31.129 17.5555 31.129 16.7223C31.129 15.4501 29.9931 14.9854 28.2211 14.9854C26.4491 14.9854 25.3131 15.4501 25.3131 16.7223C25.3131 17.5555 25.7002 18.4828 25.7002 19.3741C25.7002 20.2997 25.1071 20.9544 24.6182 21.7508C24.4904 21.0301 24.263 20.457 24.0237 19.9945C24.1171 19.8067 24.1734 19.6159 24.1734 19.4065C24.1734 18.6027 23.8172 17.8719 23.7891 16.853C23.9174 16.4766 24.0166 16.0397 24.0266 15.5172C24.4869 14.4608 25.593 13.7782 27.1611 13.5779L27.1611 5.95993C27.1611 5.47696 27.1216 5.34556 26.8308 5.14314C26.5685 4.95538 26.4615 4.79738 26.4713 4.54482L26.66 1.39853C26.6826 1.1016 26.9023 0.888101 27.1699 0.888101L29.2589 0.888101C29.5398 0.888101 29.756 1.1016 29.7786 1.39853ZM30.3451 23.4641C30.3451 23.7717 30.0925 24.0207 29.7818 24.0207L26.6439 24.0207C26.3301 24.0207 26.0806 23.7717 26.0806 23.4641C26.0806 23.1597 26.3399 22.8942 26.6439 22.8942L29.7818 22.8942C30.0925 22.8942 30.3451 23.1499 30.3451 23.4641ZM29.5865 19.0421C29.5865 19.7932 28.9752 20.4173 28.2211 20.4173C27.4603 20.4173 26.8393 19.7963 26.8393 19.0421C26.8393 18.2782 27.47 17.6572 28.2211 17.6572C28.9752 17.6572 29.5865 18.2813 29.5865 19.0421Z" fill-opacity="0.85"/><path d="M18.6502 1.10515L19.7435 2.05197C19.941 2.23351 20.0475 2.47941 20.0315 2.73507L19.7776 5.36909C19.7385 5.74905 19.6413 5.91107 19.3524 6.0269C19.1531 6.10147 19.1402 6.1503 19.1402 6.46595L19.1304 14.2513C19.1304 14.7498 19.2148 15.2058 19.4931 15.2058C19.684 15.2058 19.8921 14.9745 20.1163 14.6602C20.419 14.2314 20.8723 13.7923 21.6278 13.7923C22.5182 13.7923 23.1617 14.4661 23.1617 15.4129C23.1617 16.7024 22.2904 17.4654 22.2904 18.525C22.2904 19.4305 23.8622 20.3885 23.8622 23.0314C23.8622 26.1363 21.5594 28.1888 18.0864 28.1888C17.0588 28.1888 16.1286 28.0072 15.3326 27.665C15.5477 27.1771 15.5979 26.6029 15.4677 26.0267C16.1509 26.4343 17.0374 26.6517 18.0864 26.6517C20.7169 26.6517 22.3252 25.2846 22.3252 23.0314C22.3252 20.7067 20.7791 19.9361 20.7791 18.5281C20.7791 17.2049 21.6327 16.4219 21.6327 15.4777C21.6327 15.3724 21.5572 15.2903 21.4587 15.2903C20.9695 15.2903 20.6463 16.69 19.4447 16.69C18.8437 16.69 18.3519 16.3295 18.1526 15.5913L18.0301 15.5913C17.8405 16.326 17.3523 16.6962 16.7415 16.6962C15.5265 16.6962 15.2034 15.2903 14.7142 15.2903C14.6223 15.2903 14.5468 15.3724 14.5468 15.4777C14.5468 16.4219 15.3969 17.2049 15.3969 18.5281C15.3969 19.7849 14.1626 20.5338 13.9004 22.3365L13.0016 20.3088C13.4237 19.5233 13.8856 19.0164 13.8856 18.525C13.8856 17.4654 13.0111 16.7024 13.0111 15.4129C13.0111 14.4661 13.6614 13.7923 14.5549 13.7923C15.3037 13.7923 15.7538 14.2314 16.0566 14.6602C16.2807 14.9745 16.4889 15.2089 16.6931 15.2089C16.9586 15.2089 17.0558 14.7236 17.0558 14.2575L17.0425 6.46595C17.0425 6.15386 17.007 6.09481 16.8303 6.0269C16.5413 5.91107 16.4312 5.73929 16.3988 5.36909L16.1449 2.73507C16.1289 2.48917 16.2421 2.22685 16.433 2.05197L17.5263 1.10515C17.6816 0.969331 17.8729 0.888101 18.0864 0.888101C18.3035 0.888101 18.4948 0.969331 18.6502 1.10515ZM16.9647 18.9769C17.2493 19.0674 17.401 19.3879 17.3069 19.6724L16.3512 22.4437C16.2504 22.7282 15.934 22.8893 15.6494 22.789C15.368 22.6954 15.2131 22.3749 15.3072 22.1032L16.2661 19.3155C16.37 19.0346 16.674 18.8734 16.9647 18.9769ZM19.9103 19.3155L20.8692 22.1032C20.9633 22.3873 20.8053 22.7171 20.5177 22.7921C20.2332 22.88 19.926 22.722 19.8284 22.4437L18.8726 19.6724C18.7785 19.3755 18.9365 19.0488 19.221 18.9738C19.5117 18.8859 19.8065 19.0408 19.9103 19.3155Z" fill-opacity="0.85"/><path d="M1.0623 25.7799C0.589134 26.8528 0.97886 28.1657 2.43337 28.1657C2.93137 28.1657 3.35836 27.9686 4.01883 27.4249L7.46038 24.5631C7.6051 24.4414 7.69034 24.405 7.76891 24.405C7.84393 24.405 7.92917 24.4414 8.07079 24.5631L11.519 27.4249C12.1759 27.9686 12.6065 28.1657 13.1045 28.1657C14.5621 28.1657 14.9518 26.8528 14.4755 25.7799L8.82887 13.0311L8.82887 7.02083C8.82887 4.13421 9.32026 1.97335 9.32026 1.55122C9.32026 1.15308 9.06504 0.888101 8.67356 0.888101C8.42456 0.888101 8.2044 1.02748 7.88259 1.3879L5.07809 4.60255C4.80956 4.91636 4.87081 5.31316 5.23255 5.47827C6.34538 5.98253 6.71561 6.13968 6.71561 7.02083L6.71561 13.0076ZM2.48126 26.2545L7.68233 14.3852L7.8555 14.3852L13.0566 26.2545C13.1778 26.543 12.9447 26.6837 12.7449 26.504L8.84088 23.2515C8.42807 22.9071 8.08407 22.738 7.76891 22.738C7.45686 22.738 7.10309 22.9071 6.69028 23.2515L2.796 26.504C2.58958 26.6837 2.35651 26.543 2.48126 26.2545ZM8.99149 20.3273C9.33194 20.3273 9.62046 20.049 9.62046 19.7085C9.62046 19.3583 9.3386 19.0729 8.99149 19.0729C8.6515 19.0729 8.35987 19.3717 8.35987 19.7085C8.35987 20.0423 8.65771 20.3273 8.99149 20.3273ZM9.68616 21.955C10.0333 21.955 10.3249 21.6763 10.3249 21.3327C10.3249 20.9825 10.0333 20.6909 9.68616 20.6909C9.34616 20.6909 9.04788 20.9989 9.04788 21.3327C9.04788 21.6763 9.34616 21.955 9.68616 21.955ZM10.3835 23.5792C10.727 23.5792 11.0289 23.2875 11.0289 22.9507C11.0289 22.6071 10.7306 22.3186 10.3835 22.3186C10.0533 22.3186 9.76163 22.6169 9.76163 22.9507C9.76163 23.2942 10.0435 23.5792 10.3835 23.5792Z" fill-opacity="0.85"/>',
    active: false,
  },
]);
onMounted(async () => {
  try {
    const api = getApiBase();
    const res = await fetch(`${api}/api/albums/${route.params.id}`, {
      credentials: "include",
    });
    if (!res.ok) throw new Error("Failed to load album");
    const data = await res.json();

    console.log("Album API response:", data);
    console.log(
      "[DEBUG] onMounted isDiscogsModalOpen:",
      isDiscogsModalOpen.value,
    );
    console.log("[DEBUG] isOpen isDiscogsModalOpen:", isOpen.value);
    album.value = data.album;
    songs.value = data.album.Songs || [];
    date_added.value = formatDateTime(data.date_added);
    date_modified.value = formatDateTime(data.date_modified);
  } catch (err) {
    error.value = err.message;
  } finally {
    loading.value = false;
  }
});
</script>
