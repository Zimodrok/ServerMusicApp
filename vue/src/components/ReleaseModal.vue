<!-- <template>
  <teleport to="body">
    <div
      v-if="isOpen"
      class="fixed inset-0 z-50 flex items-center justify-center bg-black/40 backdrop-blur-sm"
    >
      <div
        class="bg-white dark:bg-zinc-900 rounded-md flex w-[900px] max-h-[700px] shadow-lg overflow-hidden relative"
      >
        <nav
          class="flex flex-col w-48 bg-zinc-100 dark:bg-zinc-800 p-3 space-y-2 overflow-y-auto"
          aria-label="Releases"
        >
          <button
            v-for="(r, i) in releases"
            :key="r.id"
            @click="selectedIndex = i"
            :class="[
              'flex space-x-3 items-center text-left rounded-md p-2 cursor-pointer transition',
              i === selectedIndex
                ? 'bg-blue-600 text-white'
                : 'hover:bg-blue-100 dark:hover:bg-zinc-700',
            ]"
          >
            <img
              :src="r.thumb"
              :alt="r.title"
              class="w-12 h-12 rounded-sm object-cover flex-shrink-0"
              loading="lazy"
            />
            <div class="flex flex-col text-sm font-semibold truncate">
              <span class="truncate">{{ r.title }}</span>
              <span
                class="font-normal text-zinc-500 dark:text-zinc-400 truncate"
                >{{ r.artist }} • {{ r.year }}</span
              >
            </div>
          </button>
        </nav>

        <section
          class="flex flex-col flex-grow border-r border-zinc-300 dark:border-zinc-700 p-4 overflow-y-auto"
        >
          <h3 class="font-semibold text-lg truncate">
            {{ selectedRelease.title }}
          </h3>
          <p class="text-zinc-600 dark:text-zinc-400 mb-3 truncate">
            {{ selectedRelease.artist }}
          </p>
          <ul class="text-sm font-mono space-y-1 select-none">
            <li
              v-for="(track, idx) in selectedRelease.tracklist || []"
              :key="idx"
              :class="[
                'flex justify-between px-2 py-0.5 rounded',
                idx === selectedTrackIndex
                  ? 'bg-blue-600 text-white font-semibold'
                  : 'text-zinc-700 dark:text-zinc-300',
              ]"
              @click="selectedTrackIndex = idx"
              tabindex="0"
            >
              <span class="shrink-0 tabular-nums">{{ track.position }}</span>
              <span class="flex-grow overflow-hidden truncate px-3">{{
                track.title
              }}</span>
              <span class="shrink-0 tabular-nums">{{ track.duration }}</span>
            </li>
          </ul>

          <div class="mt-auto pt-4">
            <label
              for="matchTracks"
              class="text-sm text-zinc-600 dark:text-zinc-400 flex items-center space-x-2"
            >
              <span>Match tracks:</span>
              <select
                id="matchTracks"
                v-model="matchMethod"
                class="form-select form-select-sm rounded border border-zinc-300 dark:border-zinc-700 bg-white dark:bg-zinc-900 text-zinc-800 dark:text-zinc-200"
              >
                <option value="auto">Automatically</option>
                <option value="manual">Manually</option>
              </select>
            </label>
          </div>
        </section>

        <aside class="flex flex-col w-72 p-4 overflow-y-auto space-y-3">
          <div class="flex space-x-3 mb-3 items-start">
            <img
              :src="selectedRelease.cover_image"
              alt="Album Cover"
              class="w-24 h-24 rounded border border-zinc-300 dark:border-zinc-700 object-cover"
              loading="lazy"
            />
            <label
              class="flex items-center space-x-2 text-sm select-none cursor-pointer"
            >
              <input
                type="checkbox"
                v-model="checkedTags.albumCover"
                class="accent-blue-600"
              />
              <span>Album Cover</span>
            </label>
          </div>

          <div class="flex flex-col space-y-1 max-h-[340px] overflow-y-auto">
            <label
              v-for="(checked, tag) in visibleTags"
              :key="tag"
              class="flex items-center space-x-2 text-sm select-none cursor-pointer"
            >
              <input
                type="checkbox"
                v-model="checkedTags[tag]"
                class="accent-blue-600"
              />
              <span
                class="capitalize text-zinc-700 dark:text-zinc-300"
                :title="formatTagLabel(tag)"
                >{{ formatTagLabel(tag) }}</span
              >
              <span
                class="flex-1 text-zinc-500 dark:text-zinc-500 truncate text-right select-text"
                >{{ tagValues[tag] || "-" }}</span
              >
            </label>

            <label
              class="flex items-center space-x-2 text-sm mt-2 pt-2 border-t border-zinc-300 dark:border-zinc-700 select-none cursor-pointer"
            >
              <input
                type="checkbox"
                v-model="checkedTags.overwrite"
                class="accent-blue-600"
              />
              <span>Overwrite</span>
            </label>
          </div>

          <button
            @click="saveTags"
            class="mt-auto bg-blue-600 text-white rounded px-4 py-2 text-sm font-medium hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            Save Tags
          </button>
        </aside>

        <button
          @click="onClose"
          aria-label="Close"
          class="absolute top-3 right-3 text-zinc-400 hover:text-zinc-600 dark:hover:text-zinc-300 transition"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-6 w-6"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            aria-hidden="true"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </button>
      </div>
    </div>
  </teleport>
</template> -->
<template>
  <teleport to="body">
    <div
      v-if="props.isOpen"
      class="fixed inset-0 z-50 flex items-center justify-center bg-black/40 backdrop-blur-sm"
    >
      <div
        class="bg-white dark:bg-zinc-900 rounded-md flex min-w-[10rem] max-w-[1500px] max-h-[700px] shadow-lg overflow-hidden relative"
      >
        <!-- Left: Releases List -->
        <nav
          class="flex flex-col max-w-1/3 bg-zinc-100 dark:bg-zinc-800 p-3 space-y-2 overflow-y-auto"
          aria-label="Releases"
        >
          <button
            v-for="(release, index) in releases"
            :key="release.id ?? index"
            @click="selectedIndex = index"
            class="flex space-x-3 items-center text-left rounded-md p-2 cursor-pointer transition"
            :class="
              index === selectedIndex
                ? 'bg-blue-600 text-white'
                : 'bg-transparent text-zinc-900 dark:text-zinc-100 hover:bg-zinc-200 dark:hover:bg-zinc-700'
            "
          >
            <img
              :src="release.thumb || release.cover_image"
              :alt="release.title"
              class="w-12 h-12 rounded-sm object-cover flex-shrink-0"
              loading="lazy"
            />
            <div class="flex flex-col text-sm font-semibold truncate">
              <span
                class="truncate"
                :title="fallbackTitle(release) || release.title"
              >
                {{ fallbackTitle(release) || release.title }}
              </span>
              <div
                class="flex items-center space-x-1 text-xs text-zinc-500 dark:text-zinc-400"
              >
                <span
                  class="truncate flex-1"
                  :title="release.artist || fallbackArtist(release)"
                >
                  {{ release.artist || fallbackArtist(release) }}
                </span>
              <span class="flex-shrink-0">
                • {{ release.year && release.year !== 0 ? release.year : "Unknown year" }}
              </span>
              </div>
            </div>
          </button>
        </nav>

        <!-- Middle: Tracklist -->
        <!-- <section
          class="flex flex-col flex-grow border-r min-w-[30rem] border-zinc-300 dark:border-zinc-700 p-4 overflow-y-auto"
        >
          <h3
            class="font-semibold text-lg truncate"
            :title="fallbackTitle(selectedRelease) || selectedRelease.title"
          >
            {{ fallbackTitle(selectedRelease) || selectedRelease.title }}
          </h3>

          <p
            class="text-zinc-600 dark:text-zinc-400 mb-3 truncate"
            :title="selectedRelease.artist || fallbackArtist(selectedRelease)"
          >
            {{ selectedRelease.artist || fallbackArtist(selectedRelease) }}
          </p>
          <ul class="text-sm font-mono space-y-1 select-none">
            <li
              class="flex justify-between px-2 py-0.5 rounded bg-blue-600 text-white font-semibold"
            >
              <span class="shrink-0 tabular-nums">A1</span>
              <span class="flex-grow overflow-hidden truncate px-3"
                >Track One Lorem</span
              >
              <span class="shrink-0 tabular-nums">4:32</span>
            </li>
            <li
              class="flex justify-between px-2 py-0.5 rounded text-zinc-700 dark:text-zinc-300"
            >
              <span class="shrink-0 tabular-nums">A2</span>
              <span class="flex-grow overflow-hidden truncate px-3"
                >Track Two Ipsum</span
              >
              <span class="shrink-0 tabular-nums">5:01</span>
            </li>
          </ul>

          <div class="mt-auto pt-4">
            <label
              class="text-sm text-zinc-600 dark:text-zinc-400 flex items-center space-x-2"
            >
              <span>Match tracks:</span>
              <select
                class="form-select form-select-sm rounded border border-zinc-300 dark:border-zinc-700 bg-white dark:bg-zinc-900 text-zinc-800 dark:text-zinc-200"
              >
                <option>Automatically</option>
                <option>Manually</option>
              </select>
            </label>
          </div>
        </section> -->

        <!-- Right: Metadata & Controls -->
        <aside
          class="flex flex-col max-w-[40rem] p-4 overflow-y-auto space-y-3"
        >
          <!-- Cover + checkbox -->
          <div
            class="flex space-x-3 mb-3 items-start font-semibold text-stone-800 dark:text-stone-400"
          >
            <img
              :src="selectedRelease.cover_image || selectedRelease.thumb"
              alt="Album Cover"
              class="w-24 h-24 rounded border border-zinc-300 dark:border-zinc-700 object-cover"
              loading="lazy"
            />
            <label
              class="flex items-center space-x-2 text-sm select-none cursor-pointer"
            >
              <input
                type="checkbox"
                v-model="tagSelection.cover"
                class="accent-blue-600"
              />
              <span>Include Album Cover</span>
            </label>
          </div>

          <!-- Fields (Genre, Subgenres, Label) -->
          <div
            class="flex flex-col space-y-1 max-h-[340px] overflow-y-auto font-semibold [&_.fontwchild]:font-normal text-stone-800 dark:text-stone-400"
          >
            <!-- Main genre -->
            <label
              class="flex items-center space-x-2 text-sm select-none cursor-pointer"
            >
              <input
                type="checkbox"
                v-model="tagSelection.genre"
                class="accent-blue-600"
              />
              <span>Genre</span>
              <span
                class="flex-1 text-zinc-500 dark:text-zinc-500 truncate text-right fontwchild"
                :title="mainGenre"
              >
                {{ mainGenre }}
              </span>
            </label>

            <!-- Subgenres / styles -->
            <label
              class="flex items-center space-x-2 text-sm select-none cursor-pointer"
            >
              <input
                type="checkbox"
                v-model="tagSelection.subgenres"
                class="accent-blue-600"
              />
              <span>Subgenres</span>
              <span
                class="flex-1 text-zinc-500 dark:text-zinc-500 truncate text-right fontwchild"
                :title="subgenresDisplay"
              >
                {{ subgenresDisplay || "—" }}
              </span>
            </label>
            <!-- Year -->
            <label
              class="flex items-center space-x-2 text-sm select-none cursor-pointer"
            >
              <input
                type="checkbox"
                v-model="tagSelection.year"
                class="accent-blue-600"
              />
              <span>Year</span>
              <span
                class="flex-1 text-zinc-500 dark:text-zinc-500 truncate text-right fontwchild"
                :title="releaseYear"
              >
                {{ releaseYear }}
              </span>
            </label>

            <!-- Comment / Description -->
            <label
              class="flex items-center space-x-2 text-sm select-none cursor-pointer"
            >
              <input
                type="checkbox"
                v-model="tagSelection.description"
                class="accent-blue-600"
              />
              <span>Description</span>
              <span
                class="flex-1 text-zinc-500 dark:text-zinc-500 truncate text-right fontwchild"
                :title="releaseDescription"
              >
                {{ releaseDescription || "—" }}
              </span>
            </label>

            <!-- Copyright / Label -->
            <label
              class="flex items-center space-x-2 text-sm select-none cursor-pointer"
            >
              <input
                type="checkbox"
                v-model="tagSelection.copyright"
                class="accent-blue-600"
              />
              <span>Copyright</span>
              <span
                class="flex-1 text-zinc-500 dark:text-zinc-500 truncate text-right fontwchild"
                :title="copyrightDisplay"
              >
                {{ copyrightDisplay || "—" }}
              </span>
            </label>
          </div>
          <button
            class="mt-auto bg-blue-600 text-white rounded px-4 py-2 text-sm font-medium hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
            @click="saveTags"
          >
            Save Tags
          </button>
        </aside>

        <button
          aria-label="Close"
          @click="onClose"
          class="absolute top-3 right-3 text-zinc-400 hover:text-zinc-600 dark:hover:text-zinc-300 transition"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-6 w-6"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            aria-hidden="true"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </button>
      </div>
    </div>
  </teleport>
</template>

<!-- <script setup>
import { ref, computed, watch, defineEmits } from "vue";
const props = defineProps({
  isOpen: Boolean,
  releases: Array,
});
const emit = defineEmits(["close", "save"]);
const releases = ref([]);
const selectedIndex = ref(0);
const selectedTrackIndex = ref(0);
const matchMethod = ref("auto");

const checkedTags = ref({
  title: true,
  trackNo: true,
  tracks: true,
  artist: true,
  albumArtist: true,
  album: true,
  year: true,
  genre: true,
  publisher: true,
  comments: true,
  releaseTime: true,
  albumCover: false,
  overwrite: true,
});

const visibleTags = computed(() => ["genre", "style", "label", "format"]);

const selectedRelease = computed(
  () => releases.value[selectedIndex.value] || {},
);

const tagValues = computed(() => {
  const r = selectedRelease.value;
  return {
    title: r.title,
    trackNo: selectedTrackIndex.value + 1,
    tracks: r.tracklist?.length || 0,
    artist: r.artist,
    albumArtist: r.artist,
    album: r.title,
    year: r.year,
    genre: Array.isArray(r.genre) ? r.genre.join(", ") : r.genre || "",
    publisher: Array.isArray(r.label) ? r.label[0] : r.label || "",
    comments: r.notes || "",
    releaseTime: r.released || "",
  };
});

function formatTagLabel(tag) {
  return tag
    .replace(/([A-Z])/g, " $1")
    .replace(/^./, (str) => str.toUpperCase());
}

function onClose() {
  this.$router.push(`/album/${this.$route.params.id}`);
}

function saveTags() {
  emit("save", {
    release: selectedRelease.value,
    selectedTrack: selectedTrackIndex.value + 1,
    tags: checkedTags.value,
    matchMethod: matchMethod.value,
  });
}
</script> -->
<script setup>
import { ref, computed, watch, defineProps, defineEmits } from "vue";

const props = defineProps({
  isOpen: Boolean,
  releases: Array,
});

const emit = defineEmits(["close", "save"]);

const selectedIndex = ref(0);
const selectedTrackIndex = ref(0);
const matchMethod = ref("auto");

watch(
  () => props.isOpen,
  (val) => {
    console.log("[DEBUG] ReleaseModal isOpen changed:", val);
    if (val) {
      selectedIndex.value = 0;
      selectedTrackIndex.value = 0;
      matchMethod.value = "auto";
    }
  },
);

const selectedRelease = computed(
  () => props.releases?.[selectedIndex.value] || {},
);

function onClose() {
  emit("close");
}
const labelDisplay = computed(() => {
  // Discogs search gives label: []string
  const labels = selectedRelease.value?.label || [];
  return labels.join(", ");
});

const releaseYear = computed(() => {
  return selectedRelease.value?.year || "";
});

// You said there is no description from search, so keep it empty for now
const releaseDescription = computed(() => {
  return "";
});

// Basic copyright guess from year + labels (you can refine later)
const copyrightDisplay = computed(() => {
  const r = selectedRelease.value;
  if (!r) return "";

  const year = r.year || "";
  const labels = (r.label || []).filter(Boolean);

  if (!labels.length && !year) return "";

  const main = labels[0] || "";
  const extras = labels.slice(1);
  const extraStr = extras.length ? " & " + extras.join(", ") : "";

  if (year && main) return `℗ ${year} ${main}${extraStr}`;
  if (main) return `℗ ${main}${extraStr}`;
  if (year) return `℗ ${year}`;
  return "";
});

const tagSelection = ref({
  cover: true,
  genre: true,
  subgenres: true,
  label: true,
  year: true,
  description: true,
  copyright: true,
});

// JS equivalent of NormalizeGenre but working from release.genre + release.style
function normalizeGenreFromRelease(release) {
  const genres = Array.isArray(release?.genre) ? release.genre : [];
  const styles = Array.isArray(release?.style) ? release.style : [];

  let main = "Unknown";
  const sub = [];
  const seen = new Set();

  if (genres.length > 0) {
    main = (genres[0] || "").trim() || "Unknown";

    // rest of genres → subgenres
    for (let i = 1; i < genres.length; i++) {
      const g = (genres[i] || "").trim();
      if (!g || g === main || seen.has(g)) continue;
      seen.add(g);
      sub.push(g);
    }

    // all styles → subgenres
    for (const s of styles) {
      const v = (s || "").trim();
      if (!v || v === main || seen.has(v)) continue;
      seen.add(v);
      sub.push(v);
    }
  } else if (styles.length > 0) {
    // no genre, fall back to first style as main
    main = (styles[0] || "").trim() || "Unknown";

    for (let i = 1; i < styles.length; i++) {
      const s = (styles[i] || "").trim();
      if (!s || s === main || seen.has(s)) continue;
      seen.add(s);
      sub.push(s);
    }
  } else {
    main = "Unknown";
  }

  return { main, sub };
}

const mainGenre = computed(() => {
  const { main } = normalizeGenreFromRelease(selectedRelease.value || {});
  return main;
});

const subgenres = computed(() => {
  const { sub } = normalizeGenreFromRelease(selectedRelease.value || {});
  return sub;
});

const subgenresDisplay = computed(() => subgenres.value.join(", "));
function fallbackArtist(release) {
  if (!release?.title) return "Unknown artist";
  // Discogs search titles are often "Artist - Album"
  return release.title.split(" - ")[0] || "Unknown artist";
}
function fallbackTitle(release) {
  if (!release?.title) return "Unknown title";
  // Discogs search titles are often "Artist - Album"
  return release.title.split(" - ")[1] || "Unknown";
}

function saveTags() {
  emit("save", {
    release: selectedRelease.value,
    selectedTrack: selectedTrackIndex.value + 1,
    matchMethod: matchMethod.value,
    tags: { ...tagSelection.value },
    normalized: {
      mainGenre: mainGenre.value,
      subgenres: subgenres.value,
      label: labelDisplay.value,
      year: releaseYear.value,
      description: releaseDescription.value,
      copyright: copyrightDisplay.value,
    },
  });
}
</script>
