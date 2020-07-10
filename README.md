#TOP DOCTORS 

>_Primero de todo tenemos que hacer uso de la función (Login_to_api) 
Esto nos permitirá poder acceder a la web de producción_

**https://www.topdoctors.es/ws/get_login/:user_crypt/:pass_crypt**

>_Una vez dentro usamos la siguinte URL para poder generar un Token_

**https://www.topdoctors.es/ws/get_doctor_data_by_id/:token/**

>_Con el Token generado, tenemos que generar un listado de los IDS de los Doctores
para realizar esto es necesario usar la siguiente URL_

**https://www.topdoctors.es/ws/get_doctor_list/:token**

>_Ahora ya tenemos un Token y los IDs de los doctores podemos acceder a la //informacion del Doctor con la siguiente URL_

**https://wwww.topdoctors.es/ws/get_doctor_data_by_id/:token/:doctor_id**

>_Esto nos generara un JSON con los siguientes datos:_
```


      	     ‘Doctor.type_doctor', = Grado máximo del doctor
             'Doctor.name', = Nombre del Doctor
             'Doctor.surname', = Apellido del doctor
             'Doctor.sexo', = Sexo del doctor (Hombre/ Mujer)
             'Doctor.profileimg',= Imagen del Doctor
             'Doctor.num_colegiado', = Id del colegio de Doctores
             'Doctor.facebook', = Perfil del Doctor en Facebook
             'Doctor.linkedin', = Perfil del Doctor en Likendin 
             'Doctor.twitter', = Perfil del Doctor en Twitter
             'Doctor.youtube', = Perfil del Doctor en Youtube
             'Doctor.vimeo', = Videos del Doctor en Vimeo
             'Doctor.instagram', = Perfil del Doctor en Instagram	
             'Doctor.pinterest', = Perfil del Doctor en Pinterest
             'Doctor.anos_experiencia', = Años del experiencia del doctor
             'Doctor.text_doctor_primer_nivel', = Descripción del perfil del Doctor 
             'Doctor.list_experiencia_professional', Experiencia del doctor
             'Doctor.cargos_organizaciones', = Cargo del Doctor
             'Doctor.logros_academicos',  = Títulos académicos del Doctor
             'Doctor.docencia', = Estudios del Doctor
             'publicacions_conferencias', = Publicaciones del Doctor	
             'premios_reconocimientos', = Premios del Doctor
             'miembro_asociaciones', = Miembro de …
             'Reviews.cantidad', = Valoraciones del Doctor	
             'Reviews.recomendar', = Recomendaciones del Doctor
             'Reviews.trato_doctor', = Valoración del trato del Doctor con los pacientes
             'Reviews.espera_consulta'  = Valoración de la espera previa a la cita
      	     'Reviews.instalaciones', = Valoración de las Instalaciones del Doctor 
             'Reviews.total', = Valoración general del Doctor*
```


