<script>
  import { onMount } from 'svelte';
  import { browser } from '$app/environment';

  let btcPrice = null;
  let priceOverMillion = 'NO';
  let displayPrice = '';
  let backgroundClass = 'bg-gray-800'; // Default background

  async function fetchBTCPrice() {
    const response = await fetch('https://api.coinbase.com/v2/prices/spot?currency=USD');
    const data = await response.json();
    btcPrice = data.data.amount;

    // Convert price to millions and truncate to 2 decimal places
    const priceInMillions = (btcPrice / 1000000).toFixed(2);
    displayPrice = `1 BTC = $${priceInMillions}M`;

    // Update display and background based on price
    if (btcPrice >= 1000000) {
      priceOverMillion = 'YES';
      backgroundClass = 'bg-green-800';
    } else {
      priceOverMillion = 'NO';
      backgroundClass = 'bg-red-800';
    }
  }

  if (browser) {
    onMount(() => {
      fetchBTCPrice();
    });
  }
</script>

<div class="flex flex-col items-center justify-center h-screen {backgroundClass} font-mono p-2">
  <div class="flex flex-col items-center justify-center flex-grow">
    <p class="text-white text-sm">
      Is the price of Bitcoin over $1M?
    </p>
    <h1 class="text-white text-9xl font-bold">
      {displayPrice ? priceOverMillion : '?'}
    </h1>
    <p class="text-white text-2xl mt-4">
      {displayPrice ? displayPrice : 'Loading...'}
    </p>
    <a href="/why" class="text-xs text-white hover:underline hover:decoration-dashed mt-4">
      Why?
    </a>
  </div>
  <footer class="py-4 bg-transparent text-center text-white text-xs sticky bottom-0">
    Built with <span class="text-red-500">&hearts;</span> by <a href="https://raghavsood.com" target="_BLANK" class="hover:underline">Raghav Sood</a>
  </footer>
</div>

<style>
  :global(html, body) {
    height: 100%;
  }
</style>

