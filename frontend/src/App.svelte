<script>
  import { onMount } from 'svelte'
  import logo from './assets/images/logo-universal.png'
  import {PreSets} from "../wailsjs/go/main/App";

  let presets = []      // will become [{ id, label }] for rendering
  let selected = ''     // bound to <select>
  let loading = false
  let error = ''

  function normalize(list) {
    return list.map((p) => {
      if (typeof p === 'string') return { id: p, name: p }
      // try common field names; tweak as needed for your real struct
      const id = JSON.stringify(p)
      const label =  String(id)
      return { id, label }
    })
  }

  async function loadPresets() {
    loading = true
    error = ''
    try {
      const data = await PreSets()
      presets = normalize(data)
      if (!selected && presets.length) selected = presets[0].id
    } catch (e) {
      error = (e && e.message) ? e.message : String(e)
    } finally {
      loading = false
    }
  }

  onMount(loadPresets)

  // defensive fallback in case your onMount timing is odd in Wails:
  if (typeof window !== 'undefined') {
    window.addEventListener('DOMContentLoaded', () => queueMicrotask(loadPresets))
  }

  function handleChange(e) {
    selected = e.target.value
    // TODO: do something with selected (emit event, call Go, etc.)
  }
</script>

<main>
  <img alt="Wails logo" id="logo" src="{logo}">
  <div class="field">
    <label for="preset">Preset</label>
    <div style="display:flex; gap:.5rem; align-items:center;">
      <select id="preset" bind:value={selected} on:change={handleChange} disabled={loading || !presets.length}>
        {#if loading}
          <option>Loading…</option>
        {:else if error}
          <option disabled>Failed to load</option>
        {:else if presets.length === 0}
          <option disabled>No presets</option>
        {:else}
          {#each presets as p (p.id)}
            <option value={p.id}>{p.label}</option>
          {/each}
        {/if}
      </select>

      <button type="button" on:click={loadPresets} disabled={loading}>
        {loading ? 'Reloading…' : 'Reload'}
      </button>
    </div>

    {#if error}
      <p style="color:#c00; margin:.25rem 0 0;">{error}</p>
    {/if}
  </div>
</main>

<style>
  .field { display: grid; gap: .4rem; }
  label { font-weight: 600; }
  select, button {
    padding: .4rem .6rem;
    border-radius: .5rem;
    border: 1px solid #7775;
    background: #1b1b1b;
    color: #eee;
  }
</style>
