package templates

const SinisterSuccess = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title></title>
  </head>
  <body>
    <table cellpadding="0" cellspacing="0" border="0" bgcolor="#ffffff" align="center">
      <tr align="center">
        <td>
          <img src="https://www.interseguro.pe/wp-content/uploads/zona-publica/siniestros/template/header.png" width="700" height="59" style="display: block;" />
        </td>
      </tr>
      <tr align="center">
        <td>
          <img src="https://www.interseguro.pe/wp-content/uploads/zona-publica/siniestros/template/banner.jpg" style="width: 700px; height: 261px; display: block;" />
        </td>
      </tr>
      <tr>
        <td height="25"></td>
      </tr>
      <tr align="center">
        <td style="font-family: Arial, Helvetica, sans-serif; font-size: 30px; font-weight: bold; color: #0855c4;">
          Hola {{.name}},
        </td>
      </tr>
      <tr>
        <td height="15"></td>
      </tr>
      <tr align="center">
        <td style="font-family: Arial, Helvetica, sans-serif; font-size: 24px; color: #454A6C;">
          Te mostramos los detalles de tu solicitud
        </td>
      </tr>
      <tr>
        <td height="30"></td>
      </tr>
      <tr>
        <td>
          <table cellpadding="0" cellspacing="0" border="0" background="https://www.interseguro.pe/wp-content/uploads/zona-publica/siniestros/template/fondo4.png" align="center" style="width: 566px; height: 479px;">
            <tr>
              <td height="30"></td>
            </tr>
            <tr align="center">
              <td>
                <table align="center" background="https://www.interseguro.pe/wp-content/uploads/zona-publica/siniestros/template/fondo2.png" align="center" style="width: 456px; height: 33px;">
                  <tr>
                    <td style="font-family: Arial, Helvetica, sans-serif; font-size: 16px; color: #0855c4; padding-left: 30px;">
                      N° de solicitud:
                    </td>
                    <td width="100"></td>
                    <td align="right" style="font-family: Arial, Helvetica, sans-serif; font-size: 16px; color: #0855c4; padding-right: 20px;">
                      {{.case}}
                    </td>
                  </tr>
                </table>
              </td>
            </tr>
            <tr align="center">
              <td>
                <table align="center" background="https://www.interseguro.pe/wp-content/uploads/zona-publica/siniestros/template/fondo2.png" align="center" style="width: 456px; height: 33px;">
                  <tr>
                    <td style="font-family: Arial, Helvetica, sans-serif; font-size: 16px; color: #0855c4; padding-left: 30px;">
                      Tipo de solicitud:
                    </td>
                    <td width="20"></td>
                    <td align="right" style="font-family: Arial, Helvetica, sans-serif; font-size: 16px; color: #0855c4; padding-right: 20px;">
                      Registro de siniestro
                    </td>
                  </tr>
                </table>
              </td>
            </tr>
            <tr align="center">
              <td>
                <table align="center" background="https://www.interseguro.pe/wp-content/uploads/zona-publica/siniestros/template/fondo2.png" align="center" style="width: 456px; height: 33px;">
                  <tr>
                    <td style="font-family: Arial, Helvetica, sans-serif; font-size: 16px; color: #0855c4; padding-left: 30px;">
                      Fecha de solicitud:
                    </td>
                    <td width="100"></td>
                    <td align="right" style="font-family: Arial, Helvetica, sans-serif; font-size: 16px; color: #0855c4; padding-right: 20px;">
                      {{.date}}
                    </td>
                  </tr>
                </table>
              </td>
            </tr>
            <tr align="center">
              <td>
                <table align="center" background="https://www.interseguro.pe/wp-content/uploads/zona-publica/siniestros/template/fondo3.png" style="width: 456px; background-repeat: no-repeat;">
                  <tr>
                    <td height="13"></td>
                  </tr>
                  <tr>
                    <td>
                      <span style="font-family: Arial, Helvetica, sans-serif; font-size: 14px; color: #454A6C; padding-left: 30px;">{{.insured_name}}</span><br />
                      <p style="font-family: Arial, Helvetica, sans-serif; font-size: 16px; color: #0855C4; margin: 0 30px; padding-bottom: 3px; border-bottom: 1px solid #C9C9C9;">{{.insured}}</p>
                    </td>
                  </tr>
                  <tr>
                    <td>
                      <span style="font-family: Arial, Helvetica, sans-serif; font-size: 14px; color: #454A6C; padding-left: 30px;">Solicitante</span><br />
                      <p style="font-family: Arial, Helvetica, sans-serif; font-size: 16px; color: #0855C4; margin: 0 30px; padding-bottom: 3px; border-bottom: 1px solid #C9C9C9;">{{.author}}</p>
                    </td>
                  </tr>
                  <tr>
                    <td height="13"></td>
                  </tr>
                </table>
              </td>
            </tr>
            <tr align="center">
              <td>
                <p style="font-family: Arial, Helvetica, sans-serif; font-size: 18px; margin: 0;">
                  <span style="color: #ffffff;">El tiempo de evaluación es de aproximadamente</span>
                  <br>
                  <span style="color: #ffffff";>{{.days}} días calendario.</span>
                </p>
              </td>
            </tr>
            <tr>
              <td height="30"></td>
            </tr>
          </table>
        </td>
      </tr>
      <tr>
        <td height="30"></td>
      </tr>
      <tr align="center">
        <td>
          <span style="font-family: Arial, Helvetica, sans-serif; font-size: 18px; color: #454A6C;">
            La respuesta será enviada al correo que registraste. También
            <br/>puedes <b>consultar el estado de tu solicitud</b>.
          </span>
        </td>
      </tr>
      <tr>
        <td height="35"></td>
      </tr>
      <tr>
        <td>
          <table align="center" background="https://www.interseguro.pe/wp-content/uploads/zona-publica/siniestros/template/fondo5.png" style="width: 524px; background-repeat: no-repeat;">
            <tr>
              <td height="15"></td>
            </tr>  
            <tr align="center">
                <td>
                  <span style="font-family: Arial, Helvetica, sans-serif; font-size: 18px; color: #0855C4;">
                    Con el seguimiento en línea <b>tendrás estos beneficios</b>
                  </span>
                </td>
            </tr>
            <tr>
              <td height="15"></td>
            </tr>  
            <tr align="center">
              <td>
                <table align="center" align="center" style="width: 456px;">
                  <tr>
                    <td style="font-family: Arial, Helvetica, sans-serif; color: #0855c4; width: 50%; padding: 0 7%; border-right: 1px dashed #8BB0FF;">
                      <div style="width: 86px; height: 86px; margin: 0 auto; display: flex; justify-content: center; align-items: center; background: #fff; border-radius: 50%">
                        <img
                          src="https://www.interseguro.pe/wp-content/uploads/zona-publica/siniestros/template/beneficio1.png"
                          style=""
                        />
                      </div>
                      <p style="font-size: 14px; font-weight: bold; margin-bottom: 3px; text-align: center;">Ahorra tiempo</p>
                      <p style="font-size: 12px; margin-top: 0; text-align: center;">Evita hacer trámites<br/>presenciales y larga<br/>espera.</p>
                    </td>
                    <td style="font-family: Arial, Helvetica, sans-serif; color: #0855c4; width: 50%; padding: 0 7%;">
                      <div style="width: 86px; height: 86px; margin: 0 auto; display: flex; justify-content: center; align-items: center; background: #fff; border-radius: 50%">
                        <img
                          src="https://www.interseguro.pe/wp-content/uploads/zona-publica/siniestros/template/beneficio2.png"
                          style=""
                        />
                      </div>
                      <p style="font-size: 14px; font-weight: bold; margin-bottom: 3px; text-align: center;">Estado de solicitud</p>
                      <p style="font-size: 12px; margin-top: 0; text-align: center;">Podrás visualizar el estado<br/>de tu solicitud desde<br/>
                        <a href="https://www.interseguro.pe/" target="_blank" style="color: #00ADED; font-style: italic;">https://www.interseguro.pe/</a></p>
                    </td>
                  </tr>
                </table>
              </td>
            </tr>
            <tr>
              <td height="15"></td>
            </tr>
            <tr align="center">
              <td>
                <a href="https://{{.subdomain}}.interseguro.pe/siniestros" target="_blank"><img src="https://www.interseguro.pe/wp-content/uploads/zona-publica/siniestros/template/boton2.png" style="width: 238px; height: 46px;" /></a>
              </td>
            </tr>
            <tr>
              <td height="30"></td>
            </tr>
          </table>
        </td>
      </tr>
      <tr>
        <td height="50"></td>
      </tr>
      <tr>
        <td>
          <hr color="#c9c9c9" />
        </td>
      </tr>
      <tr>
        <td height="30"></td>
      </tr>
      <tr align="center">
        <td>
          <span style="color: #0855c4; font-family: Arial, Helvetica, sans-serif; font-weight: bold; font-size: 18px;">Hacemos más fácil que los peruanos avancen seguros</span>
        </td>
      </tr>
      <tr>
        <td height="10"></td>
      </tr>
      <tr align="center">
        <td>
          <a href="https://www.facebook.com/interseguro/" style="text-decoration: none;">
            <img src="https://img.mailinblue.com/2581712/images/6238e5fc53294_1647896060.jpg" />
          </a>
          <a href="https://www.instagram.com/interseguroperu/?hl=es-la" style="text-decoration: none;">
            <img src="https://img.mailinblue.com/2581712/images/6238e6028d6a5_1647896066.jpg" />
          </a>
          <a href="https://www.interseguro.com.pe/blog/" style="text-decoration: none;">
            <img src="https://img.mailinblue.com/2581712/images/6238e60623502_1647896070.jpg" />
          </a>
          <a href="https://www.linkedin.com/company/interseguro-compania-de-seguros/" style="text-decoration: none;">
            <img src="https://img.mailinblue.com/2581712/images/6238e60a5ce92_1647896074.jpg" />
          </a>
        </td>
      </tr>
      <tr>
        <td height="10"></td>
      </tr>
      <tr align="center">
        <td>
          <a href="https://www.interseguro.pe/" style="color: #00ADEE; text-decoration: none; font-family: Arial, Helvetica, sans-serif; font-size: 13px;">www.interseguro.pe</a>
        </td>
      </tr>
      <tr>
        <td height="10"></td>
      </tr>
      <tr align="center">
        <td>
          <span style="color: #939393; text-decoration: none; font-family: Arial, Helvetica, sans-serif; font-size: 13px;">
            Resuelve tus dudas o consultas desde la sección de
            <a href="https://www.interseguro.pe/ayuda/contactanos" style="color: #00ADEE; text-decoration: none;">contáctanos</a>
            en nuestra web.
          </span>
        </td>
      </tr>
      <tr>
        <td height="10"></td>
      </tr>
      <tr align="center">
        <td>
          <span style="color: #939393; text-decoration: none; font-family: Arial, Helvetica, sans-serif; font-size: 10px;">Copyright © 2023 Interseguro </span>
        </td>
      </tr>
      <tr>
        <td height="50"></td>
      </tr>
    </table>
  </body>
</html>
`
