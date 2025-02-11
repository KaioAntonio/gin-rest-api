package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/KaioAntonio/gin-rest-api/controllers"
	"github.com/KaioAntonio/gin-rest-api/database"
	"github.com/KaioAntonio/gin-rest-api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	return rotas
}

func TestVerificaStatusCodeSaudacao(t *testing.T) {
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/kaio", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)

	responseMock := `{"API diz:":"E ai kaio, tudo beleza?"}`
	responseBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, responseMock, string(responseBody))
}

func TestListandoTodosAlunosHandler(t *testing.T) {
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestListandoAlunosPorCpfHandler(t *testing.T) {
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.ExibeAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678910", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorIdHandler(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.ExibeAlunoPorId)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	fmt.Println(alunoMock.Nome)
	assert.Equal(t, "Nome teste", alunoMock.Nome)
	assert.Equal(t, "12345678910", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
}

func TestDeletaAlunoHandler(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaUmAlunoHandler(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.PUT("/alunos/:id", controllers.EditaAluno)
	path := "/alunos/" + strconv.Itoa(ID)
	aluno := models.Aluno{Nome: "Nome teste", CPF: "12345678911", RG: "123456789"}
	valorJson, _ := json.Marshal(aluno)
	req, _ := http.NewRequest("PUT", path, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "12345678911", alunoMockAtualizado.CPF)
	assert.Equal(t, "123456789", alunoMockAtualizado.RG)
	assert.Equal(t, "Nome teste", alunoMockAtualizado.Nome)
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome teste", CPF: "12345678910", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}
