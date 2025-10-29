<script>
    import { onMount } from 'svelte';

    // Estado para a mensagem principal
    let dadosDaAPI = null;
    let erro = null;
    let estaCarregando = false;

    // --- NOVOS ESTADOS para os dados do cabeçalho ---
    let bitcoinPrice = null;
    let weatherData = null;
    let globalError = null; // Um erro geral para as chamadas iniciais

    // Função que busca a mensagem principal (quando o botão é clicado)
    async function buscarMensagem() {
        estaCarregando = true;
        erro = null;
        dadosDaAPI = null;

        try {
            const resposta = await fetch('http://localhost:8080/api/mensagem');
            if (!resposta.ok) throw new Error('Falha ao conectar na API.');
            dadosDaAPI = await resposta.json();
        } catch (e) {
            erro = e.message;
        } finally {
            estaCarregando = false;
        }
    }

    // --- NOVAS FUNÇÕES ---
    // onMount é executado uma vez, quando o componente é renderizado na tela.
    onMount(async () => {
        try {
            // Busca os dados em paralelo para ser mais rápido
            const [bitcoinRes, weatherRes] = await Promise.all([
                fetch('http://localhost:8080/api/bitcoin'),
                fetch('http://localhost:8080/api/clima')
            ]);

            if (bitcoinRes.ok) {
                const data = await bitcoinRes.json();
                bitcoinPrice = data.bitcoin.brl;
            }

            if (weatherRes.ok) {
                const data = await weatherRes.json();
                // Pegamos a primeira condição atual, que é a que importa
                weatherData = data.current_condition[0];
            }

        } catch (e) {
            globalError = "Não foi possível carregar os dados do cabeçalho. O backend está rodando?";
            console.error(e);
        }
    });
</script>

<header class="absolute top-0 left-0 right-0 p-4 bg-gray-800 text-white shadow-md">
    <div class="container mx-auto flex justify-between items-center text-sm">
        {#if bitcoinPrice}
            <span>🪙 Bitcoin: <strong>R$ {bitcoinPrice.toLocaleString('pt-BR', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}</strong></span>
        {:else}
            <span>🪙 Carregando...</span>
        {/if}

        {#if weatherData}
            <span>🌡️ {weatherData.temp_C}°C, {weatherData.weatherDesc[0].value}</span>
        {:else}
            <span>🌡️ Carregando...</span>
        {/if}
    </div>
</header>


<main class="flex flex-col items-center justify-center min-h-screen bg-gray-100 text-gray-800 p-4 pt-24">
    <div class="bg-white p-8 rounded-lg shadow-md max-w-lg w-full">
        <h1 class="text-3xl font-bold text-center mb-6 text-blue-600">
            Go + Svelte: API Única
        </h1>

        <div class="text-center mb-6">
            <button
                on:click={buscarMensagem}
                disabled={estaCarregando}
                class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-6 rounded-lg transition-colors disabled:bg-gray-400"
            >
                {#if estaCarregando} Buscando... {:else} Buscar Dados do Servidor {/if}
            </button>
        </div>

        {#if globalError}
            <p class="text-center text-red-500 bg-red-100 p-2 rounded">{globalError}</p>
        {/if}

        <div class="mt-4 p-4 border-t border-gray-200 text-center min-h-[150px] flex items-center justify-center">
            {#if estaCarregando}
                <p class="text-lg animate-pulse">Consultando a API Go...</p>
            {:else if erro}
                <p class="text-lg text-red-500"><strong>Erro:</strong> {erro}</p>
            {:else if dadosDaAPI}
                <div class="space-y-4 text-left">
                    <div>
                        <span class="font-semibold">Texto:</span>
                        <span class="font-mono bg-gray-100 p-1 rounded">"{dadosDaAPI.texto}"</span>
                    </div>

                    <div>
                        <span class="font-semibold">Horário da Resposta:</span>
                        <span class="text-sm">{new Date(dadosDaAPI.timestamp).toLocaleString('pt-BR')}</span>
                    </div>

                    <div class="pt-4 border-t border-dashed">
                        <p class="text-lg italic text-center text-purple-700">
                            "{dadosDaAPI.frase}"
                        </p>
                    </div>
                </div>
            {:else}
                <p class="text-gray-500">Clique no botão para ver os dados.</p>
            {/if}
        </div>
    </div>
</main>