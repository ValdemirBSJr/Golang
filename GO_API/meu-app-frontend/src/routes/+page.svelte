<script>
  let dadosDaAPI = null;
  let erro = null;
  let estaCarregando = false;

  // Voltamos a ter apenas uma função de busca
  async function buscarMensagem() {
    estaCarregando = true;
    erro = null;
    dadosDaAPI = null;

    try {
      // A chamada de API é sempre para o mesmo endpoint
      const resposta = await fetch('http://localhost:8080/api/mensagem');
      if (!resposta.ok) throw new Error('Falha ao conectar na API.');
      
      // A resposta agora contém todos os dados que precisamos
      dadosDaAPI = await resposta.json();
    } catch (e) {
      erro = e.message;
    } finally {
      estaCarregando = false;
    }
  }
</script>

<main class="flex flex-col items-center justify-center min-h-screen bg-gray-100 text-gray-800 p-4">
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
            <span class="font-semibold">Horário (João Pessoa):</span>
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