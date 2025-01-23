Um projeto over engineered para demonstrar a arquitetura que costumo usar em meus projetos:

Camadas:
    1 - Modelo
        Objetos contendo as regras de negócio (Padrão em qualquer Arquitetura)
    2 - Serviço
        Objetos contendo implementações representando ações dos usuários (equivalente aos Use Cases em outras Arquiteturas)
    3 - Aplicação
        Implementações das maneiras de como utilizar os objetos do serviço, APIs REST, GraphQL, Consumidores de fila, Desktop UIs, etc.
    4 - Infra
        Implementações dos acessos e uso de aplicações e/ou tecnologias externas à aplicação (Bancos de Dados, Mensageria, Requisições HTTP, envios SMTP, etc.)
    5 - Views
        Essas são representações para transmitir informações para visualização nas aplicações, em outras arquiteturas é o equivalente a um DTO que é retornado como resultado de uma API.
    5 - Inputs
        São os "DTOs" de entrada, por exemplo, o objeto do modelo tem o campo ID obrigatório, e o campo completed que é gerado no momento da criação do objeto, porém esses campos não são informados externamente, isso evita criar objetos de modelo/ORM com campos obrigatórios com valor null :)
    5 - Data
        A golang é tipada então aqui ficam as interfaces utilizadas nos serviços para acesso de tecnologias externas, em outros termos as implementações dessas interfaces são Adapters/Facades, ou uma "máscara" mais conveniente para o uso de bibliotecas.
        Linguagens como Python/Javascript/Ruby/Lua não precisam desse pacote

Essa estrutura é conveniente para manter as camadas importantes (Serviço e Modelo) separadas do mundo real, elas não precisam realizar tratamento de erro ou conhecer sobre detalhes de implementação de integração com outras tecnologias.

A estrutura de pastas aqui também não é importante, desde que as camadas estejam devidamente separadas, nesse exemplo a camada Data poderia muito bem ser um pacote da camada de serviço sem maiores problemas.

Para chegar nesse modelo me baseei em conceitos de arquiteturas como CleanArch, Hexagonal Arquitecture (Ports and Adapters), CQRS, MVC e MVVM. É uma arquitetura complexa demais para problemas pequenos, mas não chega a ser complexa o bastante para impedir ou atrasar o desenvolvimento de projetos menores, se adapta bem em projetos maiores ou que visam crescimento.