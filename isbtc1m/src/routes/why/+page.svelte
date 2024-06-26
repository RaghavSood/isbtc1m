<script>
  import { onMount } from 'svelte';

  onMount(() => {
    const id = 'twitter-widgets';

    // Its already there.
    if (document.getElementById(id)) {
      return;
    }

    const scriptTag = document.createElement('script');

    scriptTag.id = id;
    scriptTag.type = 'text/javascript';
    scriptTag.async = true;
    scriptTag.charset = 'utf-8';
    scriptTag.src = 'https://platform.twitter.com/widgets.js';

    document.getElementsByTagName('head')[0].appendChild(scriptTag);
  });
</script>

<div class="flex flex-col items-center justify-center bg-black text-white h-screen font-mono space-y-3 p-2">
  <div class="flex flex-col items-center justify-center flex-grow">
    <h1 class="text-4xl">Why?</h1>
    <p class="text-sm">Unit Bias is kicking in for Bitcoin. It's time to start thinking in millions.</p>
    <blockquote class="twitter-tweet tw-align-center" data-lang="en" data-dnt="true" data-theme="dark"><p lang="en" dir="ltr">The price of <a href="https://twitter.com/search?q=%24BTC&amp;src=ctag&amp;ref_src=twsrc%5Etfw">$BTC</a> is still less than $0.05M.</p>&mdash; Michael Saylor⚡️ (@saylor) <a href="https://twitter.com/saylor/status/1757132469130657905?ref_src=twsrc%5Etfw">February 12, 2024</a></blockquote>
    <a href="/" class="text-xs hover:underline hover:decoration-dashed mt-4">
      Is the price of Bitcoin over $1M?
    </a>
  </div>
  <footer class="py-4 bg-transparent text-center text-white text-xs sticky bottom-0">
    Built with <span class="text-red-500">&hearts;</span> by <a href="https://raghavsood.com" target="_BLANK" class="hover:underline">Raghav Sood</a>
</div>
