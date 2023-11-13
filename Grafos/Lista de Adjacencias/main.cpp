#include <iostream>
#include <vector>
#include <algorithm>

class Vertice{
    private:
        int num_vertice;
        std::vector<Vertice*> vertices_adjacentes;
        std::vector<int> pesos;

    public:
        Vertice(int num_vertice){
            std::vector<Vertice*> adjacentes;
            std::vector<int> pesos;

            this->num_vertice = num_vertice;
            this->vertices_adjacentes = adjacentes;
            this->pesos = pesos;
        }

        Vertice(int num_vertice, std::vector<Vertice*> vertices_adjacentes, std::vector<int> pesos){
            this->num_vertice = num_vertice;
            this->vertices_adjacentes = vertices_adjacentes;
            this->pesos = pesos;
        }

        std::vector<Vertice*> get_vertices_adjacentes(){
            return vertices_adjacentes;
        }

        std::vector<int> get_pesos(){
            return pesos;
        }

        void set_vertices_adjacentes(std::vector<Vertice*> vertices){
            this->vertices_adjacentes = vertices;
        }

        void set_pesos(std::vector<int> pesos){
            this->pesos = pesos;
        }

        bool adiciona_adjacente(Vertice* vertice, int peso){
            if (eh_adjacente(vertice)){
                return false;
            }

            vertices_adjacentes.push_back(vertice);
            pesos.push_back(peso);
            return true;
        }

        bool eh_adjacente(Vertice* vertice){
            if (std::find(vertices_adjacentes.begin(), vertices_adjacentes.end(), vertice) == vertices_adjacentes.end()){
                return false;
            }

            return true;
        }

        static Vertice* cria_vertice(int num_vertice){
            Vertice* vertice = new Vertice{num_vertice};
            return vertice;
        }

        void mostra_vertice(){
            std::cout << "Vertice " << num_vertice << "\n";
            std::cout << "\tAdjacentes: ";

            for (Vertice* vertice : vertices_adjacentes){
                std::cout << vertice->num_vertice << " ";
            }

            std::cout << "\n";
        }
};


class Grafo{
    private:
        int num_vertices;
        std::vector<Vertice*> vertices;

    public:

        Grafo(){
            this->num_vertices = 0;
        }

        int get_num_vertices(){
            return num_vertices;
        }

        std::vector<Vertice*> get_vertices(){
            return vertices;
        }

        void set_vertices(std::vector<Vertice*> vertices){
            this->vertices = vertices;
        }
        
        bool adiciona_aresta(Vertice* a, Vertice* b, int peso){
            if (!a->adiciona_adjacente(b, peso)){
                return false;
            }

            if (!b->adiciona_adjacente(a, peso)){
                return false;
            }

            return true;
        }

        // bool remove_vertice(Vertice* vertice){
        //     if (std::find(vertices.begin(), vertices.end(), vertice) == vertices.end()){

        //     }
        // }

        bool adiciona_vertice(Vertice* vertice){
            if (std::find(vertices.begin(), vertices.end(), vertice) != vertices.end()){
                return false;
            }

            num_vertices++;
            vertices.push_back(vertice);

            return true;
        }
};


int main(){
    Grafo grafo;
    Vertice* vertices[5];

    for (int i = 0; i < 5; i++){
        vertices[i] = Vertice::cria_vertice(i);
        grafo.adiciona_vertice(vertices[i]);
    }


    grafo.adiciona_aresta(vertices[0], vertices[1], 2);
    grafo.adiciona_aresta(vertices[0], vertices[2], 5);
    grafo.adiciona_aresta(vertices[1], vertices[3], 3);
    grafo.adiciona_aresta(vertices[3], vertices[2], 6);
    grafo.adiciona_aresta(vertices[2], vertices[4], 4);

    for (Vertice* vertice : vertices){
        vertice->mostra_vertice();
    }   

}