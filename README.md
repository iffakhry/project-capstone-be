<div id="top"></div>

<!-- PROJECT LOGO -->
<br/>
<div align="center">
<!--  mengarah ke repo  -->
  <a href="https://github.com/sahrilmahendra/project2-airbnb">
    <img src="images/Logo.png" width="267" height="80">
  </a>

  <h3 align="center">Barengin</h3>

  <p align="center">
    Final Project Capstone Program Immersive Alterra Academy
    <br />
    <a href="https://app.swaggerhub.com/apis-docs/supriadi15001/final-project_alta_barengin/1.0"><strong>Explore the docs Open API »</strong></a>
    <br />
  </p>
</div>


<!-- ABOUT THE PROJECT -->
## About The Project

<!-- [![Product Name Screen Shot][product-screenshot]](https://example.com) -->

Barengin merupakan suatu platform yang menjadi wadah untuk mempertemukan calon customer yang ingin berlangganan produk digital tertentu dengan calon customer lain untuk mendapatkan keuntungan layanan premium secara patungan.

Berikut fitur yang terdapat dalam barengin :
| Feature | Admin | Customer | Guest
|:----------| :----------:| :----------:|:----------:|
| Signup | No | No | Yes
| Login | Yes | Yes | No
|---|---|---|---|
| Get all users | Yes | No | No
| Get user by id | Yes | Yes | No
| Update user by id | Yes | Yes | No
| Delete user by id | Yes | Yes | No
|---|---|---|---|
| Create product | Yes | No | No
| Get all products | Yes | Yes | Yes
| Get product by id | Yes | Yes | Yes
| Update product by id | Yes | No | No
| Delete product by id | Yes | No | No
|---|---|---|---|
| Create group product | No | Yes | No
| Get all group product | Yes | Yes | Yes
| Get group product by id group product | Yes | Yes | Yes
| Get group product by id product | Yes | Yes | Yes
| Get group product by status | Yes | Yes | Yes
|---|---|---|---|
| Create order | No | Yes | No
| Get all orders by id group | Yes | No | No
| Get all order by id user | Yes | Yes | No
| Get order by id order | Yes | Yes | No
| Update order by id order | Yes | No | No
|---|---|---|---|

<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

<!-- * [Golang](https://golang.org/)
* [Echo Framework](https://echo.labstack.com/)
* [MySQL](https://www.mysql.com/)
* [Gorm](https://gorm.io/)
* [JWT](https://echo.labstack.com/cookbook/jwt)
* [Assert](https://pkg.go.dev/github.com/stretchr/testify/assert)
* [VS Code](https://code.visualstudio.com/) -->
![VS Code](https://img.shields.io/badge/-Visual%20Studio%20Code-05122A?style=flat&logo=visual-studio-code&logoColor=FFFFFF)&nbsp;
![MySQL](https://img.shields.io/badge/-MySQL-05122A?style=flat&logo=mysql&logoColor=FFFFFF)&nbsp;
![Golang](https://img.shields.io/badge/-Golang-05122A?style=flat&logo=go&logoColor=FFFFFF)&nbsp;

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- How to Use -->
## How to Use
### * Running on Local Server
- Install Golang, Postman, MySQL Management System (ex. MySQL Workbench)
- Clone repository with HTTPS:
```
git clone https://github.com/project-capstone/project-capstone-be.git
```
* Create File `.env`:
```
export CONNECTION_DB="[username]:[password]@tcp(127.0.0.1:3306)/[db_name]?charset=utf8&parseTime=True&loc=Local"

export GOOGLE_API_KEY="[geocode_api_key]"

export KEY_XENDIT = "[xendit_key]"
```
* Run `main.go` on local terminal
```
$ go run main.go
```
* Run the endpoint according to the OpenAPI Documentation (Swagger) via Postman 

<br/>

### * Running on Cloud Server
- Create new instance
- Configure inbound rules, add http, https, mysql ports
- Create RDS Instance (MySQL)
- Run instance
- Install docker in instance
```
sudo apt install docker.io
```
- Change docker access mode
```
sudo chmod 777 /var/run/docker.sock
```
- Add `ssh key` in instance
```
ssh-keygen
```
- Add  `ssh key` in github account, ssh & gpg key -> new ssh -> fill ssh from `id_rsa.pub` value
- Clone repository with SSH:
```
git@github.com:project-capstone/project-capstone-be.git
```
- [optional] rename project name
```
mv [project_name] [new_project_name ex: app]
```
- Create File `.env` in directory project on cloud server:
```
export CONNECTION_DB="[username]:[password]@tcp([rds_instance_url]:3306)/[db_name]?charset=utf8&parseTime=True&loc=Local"

export GOOGLE_API_KEY="[geocode_api_key]"

export KEY_XENDIT = "[xendit_key]"
```
- Configure CI/CD in local
- Push CI/CD to cloud server
- Run the endpoint according to the OpenAPI Documentation (Swagger) via Postman
<!-- ERD -->
## ERD
<img src="images/erd.png">
<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Project Link : [https://github.com/project-capstone](https://github.com/project-capstone)<br/>
Open API Documentation : [https://app.swaggerhub.com/apis-docs/supriadi15001/final-project_alta_barengin/1.0#/](https://github.com/project-capstone)&nbsp;
<!-- :heart: -->
<!-- CONTRIBUTOR -->
Contributor :
<br>
[![Linkedin: Sahril Mahendra](https://img.shields.io/badge/-SahrilMahendra-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/sahril-mahendra/)](https://www.linkedin.com/in/sahril-mahendra/)
[![GitHub Sahril Mahendra](https://img.shields.io/github/followers/sahrilmahendra?label=follow&style=social)](https://github.com/sahrilmahendra)

[![Linkedin: Nuril H](https://img.shields.io/badge/-NurilH-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/sahril-mahendra/)](https://www.linkedin.com/)
[![GitHub Nuril H](https://img.shields.io/github/followers/NurilH?label=follow&style=social)](https://github.com/NurilH)

[![Linkedin: Supriadi](https://img.shields.io/badge/-Supriadi-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/sahril-mahendra/)](https://www.linkedin.com/)
[![GitHub S](https://img.shields.io/github/followers/sprdx?label=follow&style=social)](https://github.com/sprdx)
<br>
Mentor :
<br>
<!-- https://www.linkedin.com/in/iffakhry/ -->
[![Linkedin: Fakhry Ihsan](https://img.shields.io/badge/-FakhryIhsan-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/iffakhry/)](https://www.linkedin.com/in/iffakhry/)
[![GitHub Fakhry Ihsan](https://img.shields.io/github/followers/iffakhry?label=follow&style=social)](https://github.com/iffakhry)


<p align="right">(<a href="#top">back to top</a>)</p>
<h3>
<p align="center">:copyright: 2021 | Built with :heart: from us</p>
</h3>