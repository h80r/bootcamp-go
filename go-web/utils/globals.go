package utils

import "go-web/models"

// 2. Você deve ter um array da entidade na memória (no nível global), no qual todas as
// requisições que são feitas devem ser salvas.
var database *[]models.User
