package converter

import (
	"testing"

	"fmt"
	"os"
	"time"

	"runtime"
)

var htmlPage = `<!DOCTYPE html>
<html>
  <head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
  <title>Hailo Driver Payment Report</title>
  <style>
    /* RESET */
    /* ----------------------------------------- */

    /**
     * Eric Meyer"s Reset Stylesheet
     *
     * v2.0
     * 2011-01-26
     * Author: Eric Meyer - http://meyerweb.com/eric/tools/css/reset/
     */

    html, body, div, span, applet, object, iframe,
    h1, h2, h3, h4, h5, h6, p, blockquote, pre,
    a, abbr, acronym, address, big, cite, code,
    del, dfn, em, img, ins, kbd, q, s, samp,
    small, strike, strong, sub, sup, tt, var,
    b, u, i, center,
    dl, dt, dd, ol, ul, li,
    fieldset, form, label, legend,
    table, caption, tbody, tfoot, thead, tr, th, td,
    article, aside, canvas, details, embed,
    figure, figcaption, footer, header, hgroup,
    menu, nav, output, ruby, section, summary,
    time, mark, audio, video {
      margin: 0;
      padding: 0;
      border: 0;
      font-size: 100%;
      font: inherit;
      vertical-align: baseline;
    }
    /* HTML5 display-role reset for older browsers */
    article, aside, details, figcaption, figure,
    footer, header, hgroup, menu, nav, section {
      display: block;
    }
    body {
      line-height: 1;
    }
    ol, ul {
      list-style: none;
    }
    blockquote, q {
      quotes: none;
    }
    blockquote:before, blockquote:after,
    q:before, q:after {
      content: "";
      content: none;
    }
    table {
      border-collapse: collapse;
      border-spacing: 0;
    }

    /* <styles> */

    /* <print rules> */

    thead {
      display: table-header-group;
    }

    tfoot {
      display: table-footer-group;
    }

    tr {
      page-break-inside:avoid;
      page-break-after:auto;
    }

    td {
      page-break-inside:avoid;
      page-break-after:auto;
    }

    /* </print rules> */

    body {
      background: #fffffe;
      font-family: Arial, sans-serif;
      font-size: 12px;
    }

    .clear {
      clear: both;
    }

    a {
      text-decoration: none;
    }

    #wrapper {
      width: 1280px;
      margin: 15px auto;
      position: relative;
    }

    #logo {
      width: 78px;
      float: left;
      margin-right: 15px;
      margin-top: 2px;
    }

    #header {
      padding-bottom: 26px;
    }

    .header-inner {
      float: left;
      width: 650px;
    }

    h1 {
      font-size: 22px;
      font-weight: bold;
      text-transform: uppercase;
      margin-bottom: 3px;
    }

    .driver-details {
      color: #8a8b8a;
      font-size: 18px;
    }

    /* <heading table> */

    .statement-number {
      color: #746c4c;
      font-size: 16px;
      margin-top: 8px;
    }

    .date {
      color: #c1b68b;
      font-size: 15px;
      margin-top: 25px;
    }

    /* </heading table> */

    table {
      width: 100%;
    }

    table.align-top td {
      vertical-align: top;
    }
    
    table.align-top td:first-child {
      width: 50%;
    }
    
    td.td-empty-space {
      width: 10px;
    }
    
    /* <tables shared styles> */

    .total,
    .jobs,
    .factura,
    .payment-details,
    .retried-payments,
    .summary,
    .key-stats {
      border: 2px solid #ada185;
    }

    .summary caption,
    .payment-details caption,
    .retried-payments caption,
    .factura caption,
    .jobs caption,
    .total caption {
      padding: 10px 15px;
      background: #FEF3D5;
      color: #66593c;
      text-align: left;
      font-size: 20px;
      font-weight: bold;
    }

    .jobs caption,
    .factura caption,
    .total caption {
      padding: 15px;
    }

    .summary caption,
    .payment-details caption,
    .retried-payments caption,
    .jobs caption {
      border-bottom: 2px solid #ada185;
    }
    
    .alignright {
      text-align: right;
    }

    /* </tables shared styles> */

    /* <key stats table> */

    .key-stats {
      float: right;
      width: 380px;
    }

    .key-stats caption {
      padding: 5px 10px;
      background: #FBBE00;
      color: #000;
      text-align: left;
      font-size: 16px;
      font-weight: bold;
    }

    .key-stats th,
    .key-stats td {
      padding: 5px 10px;
    }

    .key-stats th {
      font-size: 15px;
    }

    .key-stats td {
      font-size: 14px;
    }

    .key-stats thead {
      background: #f3efde;
    }

    .key-stats strong {
      font-weight: bold;
    }

    .key-stats .deemphasized {
      padding-top: 2px;
    }

    .key-stats .deemphasized span {
      color: #B1B2B1;
      font-size: 12px;
      line-height: 16px;
    }

    .key-stats tbody strong {
      display: inline-block;
      padding-left: 10px;
    }

    /* </key stats table> */

    /* <total table> */

    .total {
      margin-bottom: 25px;
      font-size: 13px;
    }

    .total caption {
      padding: 15px;
      background: #FEF3D5;
      color: #66593c;
      text-align: left;
      font-size: 20px;
      font-weight: bold;
    }

    .total table td {

    }

    .total table td {

    }

    .total table td,
    .total table th {

    }

    .total table td:first-child,
    .total table th:first-child {
      border-left: none;
    }
    .total thead th {
      background: #f5f6f6;
      color: #94846a;
      font-weight: bold;
      font-size: 14px;
      vertical-align: middle;
    }

    .total td span,
    .total th span,
    .total th strong {
      display: block;
      text-align: right;
    }

    .total th span {
      padding: 5px 13px;
    }

    .total td span {
      padding: 4px 13px;
      line-height: 18px;
    }

    .total th strong {
      color: #66593c;
      font-size: 18px;
      padding: 3px 0 0;
    }

    .total thead .calculation span {
      font-size: 16px;
      text-align: left;
    }

    .total td:first-child span {
      text-align: left;
    }

    .total em {
      color: #525025;
      display: block;
      font-size: 12px;
      line-height: 18px;
      margin-right: 5px;
      padding: 5px 8px;
      background: #F3EFDE;
      -moz-border-radius: 2px;
      -webkit-border-radius: 2px;
      -o-border-radius: 2px;
      -ms-border-radius: 2px;
      border-radius: 2px;
    }

    .total .adjustment em {
      text-align: left;
    }

    .total .adjustment em {
      /*max-width: 150px;
      min-width: 150px;*/
    }

    .total tr:first-child span {
      padding-top: 10px;
    }

    .total thead span {
      padding-bottom: 10px;
    }

    .total td:first-child span {
      /*min-width: 180px;*/
    }

    .total .deemphasized span {
      color: #B1B2B1;
      font-size: 13px;
      line-height: 18px;
    }

    .total td.deemphasized span {
      padding-top: 8px;
      padding-bottom: 10px;
    }

    /* <plus, munus icons> */

    .icon {
      width: 27px;
      height: 27px;
      padding: 0 !important;
      position: relative;
      display: inline-block;
      float: left;
    }

    .total td .icon {
      display: inline-block;
    }

    .horizontal-bar,
    .vertical-bar {
      background: #d4d4d4;
      padding: 0 !important;
      display: block;
    }

    .horizontal-bar {
      height: 4px;
      left: 5px;
      line-height: 30px;
      position: absolute;
      top: 12px;
      width: 18px;
    }

    .plus .vertical-bar {
      height: 18px;
      left: 12px;
      line-height: 30px;
      position: absolute;
      top: 5px;
      width: 4px;
    }

    .icon .horizontal-bar.top {
      top: 9px;
    }

    .icon .horizontal-bar.bottom {
      top: 15px;
    }

    .total .cash-collected {
    }

    .total .cash-collected span {
      padding-left: 0;
      padding-right: 0;
    }

    .total tr .adjustment span {
      padding-left: 0;
    }

    .total tr .adjustment span {
      padding-left: 18px;
    }

    .total tr .charges span,
    .total tr .charge span {
      padding-left: 0;
    }

    .total thead tr {
      border-bottom: 12px solid #ffffff;
    }

    .total tbody tr:first-child span {
      padding-top: 0;
    }

    .total .td-charge-percent {
      border-left: 1px solid #C1B68B;
    }

    .total .td-charge-percent span {
      padding-right: 0;
      padding-left: 18px;
    }

    .total .td-charge,
    .total .td-payment,
    .total .td-adjustment
     {
      border-right: 1px solid #C1B68B;
    }

    .total .td-adjustment span,
    .total .th-adjustment span {
      padding-right: 30px;
    }

    .total .td-total,
    .total .th-total {
      padding-right: 20px;
    }

    .total .td-charge,
    .total .th-charge {
      padding-left: 5px;
      padding-right: 12px;
    }

    .total .adjustment {
      border-right: 1px solid #C1B68B;
      padding-right: 8px;
    }

    .total .tr-total .td-charge,
    .total .tr-total .td-charge-percent,
    .total .tr-total .td-payment,
    .total .tr-total .td-adjustment,
    .total .tr-total .td-total {
      border: none;
    }

    .total .tr-total td {
      background: #f5f6f6;
    }

    .total .tr-total .td-adjustment .icon {
      margin-left: 35px;
    }

    .total .td-adjustment-label span {
      text-align: left;
      padding-left: 18px;
      padding-right: 0;
    }

    .total .tr-total td:first-child span {
       color: #94846A;
      font-size: 14px;
      font-weight: bold;
    }

    #wrapper .total table .tr-total td {
      line-height: 28px;
      padding-top: 5px;
      padding-bottom: 5px;
      border-top: 12px solid #ffffff;
    }

    .total .tr-total td span {
      color: #66593C;
      font-size: 18px;
      font-weight: bold;
    }

    .total td {
      vertical-align: middle;
    }

    .total .tr-total .outer {
      min-height: 28px;
      line-height: 28px;
    }

    .total .outer-equals span {
      padding: 0;
    }

    .total .arrow-outer {
      text-align: center;
      margin-bottom: 10px;
    }

    .total .arrow {
      text-align: center;
      position: relative;
      width: 14px;
      margin: 0 auto;
    }

    .total .bar {
      background: #C1B68B;
      display: block;
      min-height: 40px;
      margin: 0 auto;
      padding: 0 !important;
      width: 3px;
    }

    .total .triangle {
      border-left: 7px solid #FFFFFF;
      border-right: 7px solid transparent;
      border-top: 14px solid #C1B68B;
      height: 0;
      position: absolute;
      width: 0;
      left: 0;
    }

    /* </plus, munus icons> */

    /* </total table> */

    /* <jobs table> */

     .jobs table td,
	.factura table td {
      border-top: 2px solid #e1e1e1;
    }

    .jobs table td,
    .jobs table th {
      border-left: 2px solid #e1e1e1;
    }
	
    .factura table td,
    .factura table th {
      border-left: 2px solid #e1e1e1;
    }

    .jobs table td:first-child,
    .jobs table th:first-child {
      border-left: none;
    }
    .jobs thead th {
      background: #f5f6f6;
      color: #444444;
      font-weight: bold;
      font-size: 13px;
      vertical-align: middle;
    }

    .factura table td:first-child,
    .factura table th:first-child {
      border-left: none;
    }
    .factura thead th {
      background: #f5f6f6;
      color: #444444;
      font-weight: bold;
      font-size: 13px;
      vertical-align: middle;
    }
	
    .jobs td span,
    .jobs th strong {
      display: block;
    }

    .jobs td span {
      padding: 12px 8px 10px;
    }

    .jobs th strong {
      padding: 10px 15px;
    }
	
    .factura td span,
    .factura th strong {
      display: block;
    }

    .factura td span {
      padding: 12px 13px 10px;
    }

    .factura th strong{
      padding: 5px 15px;
    }
	
	.factura td strong {
      font-weight: bold;
	}

    /* </jobs table> */

    .hr {
      border-bottom: 4px dashed #dddddd;
      line-height: 30px;
      margin: 60px 0 30px;
      text-align: center;
    }

    .hr h2 {
      background: #FFFFFF;
      bottom: -18px;
      color: #94846a;
      display: inline-block;
      font-size: 15px;
      font-weight: bold;
      line-height: 30px;
      padding: 0 35px;
      position: relative;
      text-align: center;
    }

    /* <payment details> */

    .payment-details table {
      width: 100%;
    }

    .payment-details .heading {
      color: #f9b700;
      font-weight: bold;
      padding-right: 10px;
    }

    .payment-details td {
      font-size: 15px;
      vertical-align: middle;
    }

    .payment-details td:first-child {
      width: 54%;
    }

    .payment-details span {
      padding: 9px 5px 0 5px;
      display: block;
    }

    .payment-details .deemphasized {
      color: #b1b2b1;
      font-size: 11px;
    }

    .payment-details .note .deemphasized {
      padding-bottom: 12px;
    }

    /* </payment details> */

    /* <summary> */

    .summary {
      background: #fcfcfc;
    }

    .summary thead th {
      font-weight: bold;
      font-size: 14px;
      vertical-align: middle;
      text-align: right;
    }

    .summary tr td:first-child {
      width: 230px;
      font-weight: bold;
      font-size: 13px;
    }

    .summary span,
    .summary strong {
      display: block;
    }

    .summary span {
      padding: 2px 12px 5px;
    }

    .summary th strong {
      padding: 6px 14px;
    }

    .summary strong {
      font-weight: bold;
    }

    /* </summary> */

    .number {
      text-align: right;
    }

    th {
      text-align: left;
    }

    /* <extra> */

    .extra {
      margin-top: 15px;
    }

    .extra h3 {
      font-size: 15px;
      margin-bottom: 10px;
    }

    .extra p {
      color: #B1B2B1;
    }

    .extra a {
      color: #F9B700;
    }
    p.pagebreakhere {
      page-break-before: always
    }
    /* </extra> */
  </style>
</head>
<body>
<div id="wrapper">
  <div id="header">
    <img id="logo" title="Hailocab.com" alt="Hailocab.com" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEgAAABMCAYAAADOfPFRAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAABT5JREFUeNrsXG1kVXEYf3abMS7LZYyxjLIsmxERyyZK9CFLWdqXZTSmLM1Y+tSUjZQiUorI0hgjxRRpGlMa05jSSJcxYhpj3Fzr/3SfY2dnz3l5zrn3npd7fjwfdt7uPb/7/J/3/8r+vgdEUsl52I51JS/AHBeZY8+UZMAaTUoOm5x7rmTD4t4KJd3M8ceGv1NKzhiOrSqZACHKiKA6Jb8M59JK9ljcu8kc261kzeYzx5gfQ0O7kmmLe6uU/OHeg/kRvhqOLShplhKUgOKj1eJcGwQMxSaonrQ1JsglAWibdpUyQUdtzlcqORRr0HbDKbFRkSaojrE/T8n9BtYOJXxeXjMkgbVDCR+XFwai80pmmSCvMdagHDFZRoMCtcwSNjZj00K82h9Nc+aY9KKt1DSozcT+AJEzF1RP5hdBWYPtMdqhGiUNpUSQ0f7Mk5E2alPgtKjc4pybcoeZ/ak3WV5WBLVRnBRYgjCA680DQa0mzz5oOJY2GPLAa1AhA8QbJE4y/3TUbZAXl+27FhWaoFole2OCCqM9gQgYy30g6LaSJZPre2B7PQhzsmolv0uFIAwQR5kSh4Yk7CyY4TKbzNNyf2Tjta8Vk6BaJhr+bEEOF1FrJOeDoJRNaJLmCEoUeXlN29zDJa6+GupCEsS92DubezKkZXq0QK4f5qsNyjAZ9YqDXxsYG6OhynBN1mQJGTFJtkiPBiIua/K5Rmw4vM72fbXOagyfs/mYoJigmKAYhYqkWyghTek8TrKI77AAWzNJc5TGrPlJEBa6Tio5BgEcNiAgSeif31DsteH2QU7dfAUlkhdJY8IETG1wcu0uuCi+OSEIteUeeKvrBAGoRfeVDEs0ys5Ijyh5HQFyEDhaM6TkAyXSnghKEjFDEXRMWE75Ag7nkBIm9uYtLa2oAhuTH53YU46gm2A+phsloCK8tAtJjASdUDJYQnEgxmxPnBJUaXdxRNFJimFLUI/EukcM/U7ioG+Qn4kKfehfrGXiNbXBQtx+YLot5TrX54YcHHDAfRJTkKs3Z3zUgkbyvN0gH+HDdOmSkitmS6zDxRfCjSG4l2OA8p2Mz8tkEXI9twNKLrv4Ph1WNkjq1p8pOQvWLRw/8UDJKdheI7cDN6bzn6BdIJtux3XaFwLDO0W5l7RSsYOgOnLxTjHqpXxQZNwSLrVGjiCJa8+Ci01pPpc6poTLbAdBNUJDuAbhgqQ/VsMRJOlaLoUwCFwUXFvFEVQheMBaCAlaF1xb6SSbjxqyXm5OCB+QDCFBEhOywREkWTZ1ISSo3itBkvG2JmHMFARIguBljqCfQiPWGbLlJSkdpzmC0kJLPyj0fH7iqlDjZ82M9LwwHB8JATk44Sbtynw2c/MzLn6ZJwH2amgG3go1/TtXndAIeuPiS2CJFv/fx3UuC/YBmFN2K/kEuW6F1JmwPWZ9yfUHeO+g5nWyQgDsb3kdojjCrST9dMdDJXc8fkhYW9QLZmZGn2pgbXkFShPDVqmGPqnrLUFycOx4wglBiFeQq+eWCjBIvmCXrBoxIIyLwgosxZ6zcyoJkxuPg3xSPUzApPQ0Fxg6IQgogW2HcNWfJcuq3WnsZ1UwQ6ONva8uLssNIbIUyjQ70RwnBGnAvfP7INeWDWNNep2Iwd57nzAxd7WZpZVKCChNASVlhQK/cci1xa0MMfbk++n6FXon7MriMEeX190+FZSHNenqL8WOpldhq6a1TB5YahLGSLswr6wn+4T7/ce9TtpnqIYyC9FAikwKlnNwe2Z1OcRAPCXXP05ZvbbBMPVPgAEAtP8QWb3b/GcAAAAASUVORK5CYII=" >


 







  


<!--

TITLE

 -->

<div class="header-inner">
    <h1></h1>
    <h1>Estado de Cuenta Semanal - Barcelona</h1>
    <p class="driver-details">Conductor: Evgeny Syrtsov (53775)</p>
    <p class="driver-details">DNI/NIE/NIF: x5307831Y</p>

    <p class="date">
    
    Lunes 26 de Octubre 2015 a Domingo 1 de Noviembre 2015
    
    </p>
</div><!-- end header-inner -->

<div class="clear"></div>

</div><!-- end header -->

<!--

PAYMENT

 -->



<div class="total">
  <table>
    <caption>Pago al conductor &nbsp;&nbsp;&nbsp; 35,40 €</caption>
    <thead>
      <tr>
        <th valign="middle" class="calculation" style="width:200px"><span>Tipo de servicio</span></th>
        <th valign="middle" class="th-total"><span>Total</span></th>
        <th valign="middle" class="th-charge" colspan="2"><span>Coste Hailo (IVA incl.)</span></th>

        <th valign="middle" class="th-adjustment" colspan="2"><span>Ajustes (IVA incl.)</span></th>


        <th valign="middle" class="cash-collected" colspan="2"><span>Efectivo ya recibido<br />por el conductor</span></th>


        <th valign="middle" colspan="2"><span>Pago Hailo<br />al conductor</span></th>
      </tr>
    </thead>
    <tfoot>
      <tr>
        <td colspan="8" class="deemphasized">
          <span>El coste Hailo cubre todos los costes de procesamiento de tarjetas de crédito y los posibles retrocesos <br />*Incluye las correciones a servicios de semanas anteriores - más detalles a continuación</span>
        </td>
      </tr>
    </tfoot>



   <tbody>
        <tr>

         
            <td valign="middle"><span>Servicios Hailo - Efectivo</span></td>
            <td valign="middle" class="td-total"><span>38,00 €</span></td>
            <td valign="middle" class="td-charge-percent">
              
                <span><em>10,00% + IVA</em></span>

            </td>
            <td valign="middle" class="td-charge"><span>-4,60 €</span></td>
         
         
         
        
            <td valign="middle" class="td-adjustment-label"><span><em>Pagos por anulaciones</em></span></td>
            <td valign="middle" class="td-adjustment"><span>0,00 €</span></td>
        

        
            <td valign="middle"></td>
            <td valign="middle" rowspan="3">
              <div class="arrow-outer">
                <div class="arrow">
                  <div class="bar">&nbsp;</div>
                  <div class="triangle"></div>
                 </div>
              </div>
            </td>

          <td valign="middle"></td>
          <td valign="middle" rowspan="3 "><div class="arrow-outer"><div class="arrow"><div class="bar">&nbsp;</div><div class="triangle"></div></div></div></td>
        </tr>
        
        
        <!-- We don"t display this if there are no data -->
        <tr>
        
          <td valign="middle"><span>Servicios Hailo - Cuenta y Tarjeta</span></td>
          <td valign="middle" class="td-total"><span>45,50 €</span></td>
          <td valign="middle" class="td-charge-percent">
          
            <span><em>10,00% + IVA</em></span>
          
          </td>
          <td valign="middle" class="td-charge"><span>-5,51 €</span></td>

          
            <td valign="middle" class="td-adjustment-label"><span><em>Otros ajustes*</em></span></td>
            <td valign="middle" class="td-adjustment"><span>0,00 €</span></td>
          

          <td valign="middle"></td>

          <td valign="middle"></td>
        
        </tr>
        
        


          <tr>
          
            <td valign="middle"><span>Carreras no Hailo - Tarjeta</span></td>

            <td valign="middle" class="td-total">
            <span>
              
            </span>
            </td>

            <td valign="middle" class="td-charge-percent">
                
                <span>
                    <em>0,00% + IVA</em>
                </span>
                

          </td>
            <td valign="middle" class="td-charge">
              
                <span>0,00 €</span>
              
            </td>

            
              <td valign="middle" class="td-adjustment-label"><span><em>Códigos Hailo</em></span></td>
              <td valign="middle" class="td-adjustment"><span>0,00 €</span></td>
            

            <td valign="middle"></td>
            
          </tr>


        
          <tr>
            <td valign="middle"><span>Pagar con Hailo</span></td>
            <td valign="middle" class="td-total">
            <span>0,00 €</span></td>
            <td valign="middle" class="td-charge-percent"></td>
            <td valign="middle" class="td-charge"><span>0,00 €</span></td>
            <td valign="middle" class="td-adjustment-label"></td>
            <td valign="middle" class="td-adjustment"></td>
            <td valign="middle"></td>
          </tr>
        

        
        <!-- TOTAL LINE -->
        
        <tr class="tr-total">
          <td valign="middle"><span>Total</span></td>
          <td valign="middle" class="td-total"><span>83,51 €</span></td>
          <td valign="middle" class="td-charge" colspan="2"><span class="outer"><span class="icon minus"><span class="horizontal-bar"></span><span class="vertical-bar"></span></span>10,11 €<div class="clear"></div></span></td>

          <td valign="middle" class="td-adjustment" colspan="2"><span class="outer"><span class="icon plus"><span class="horizontal-bar"></span><span class="vertical-bar"></span></span>0,00 €<div class="clear"></div></span></td>

            <td valign="middle"><span class="outer"><span class="icon minus"><span class="horizontal-bar"></span><span class="vertical-bar"></span></span></span></td>
            <td valign="middle"><span class="outer">38,00 €<div class="clear"></div></span></td>

          <td valign="middle" class="outer outer-equals"><span><span class="icon equals"><span class="horizontal-bar top"></span><span class="horizontal-bar bottom"></span></span><div class="clear"></div></span></td>
          <td valign="middle" class="outer"><span>35,40 €<div class="clear"></div></span></td>
        </tr>
        
        <!-- END TOTAL LINE -->
        
    </tbody>
  </table>
</div>




<!--
JOBS
 -->


<div class="jobs">
  <table>
    <caption>Servicios</caption>
    <thead>
      <tr>
        

        <th valign="middle"><strong>Fecha</strong></th>
        <th valign="middle"><strong>Ref.<br />Servicio</strong></th>
        <th valign="middle"><strong>Tipo de<br />servicio</strong></th>
        <th valign="middle" style="width:20px"><strong>Hora<br />(UTC)</strong></th>

        <th valign="middle" style="font-size:12px;width:120px"><strong>Lugar de PAB -<br />Lugar de finalización</strong></th>
        
        
        <th valign="middle" class="alignright" style="font-size:12px"><strong>Importe total + propina</strong></th>
        

        <th valign="middle" class="alignright"><strong>Coste <br />Hailo bruto</strong></th>

        
          <th valign="middle" class="alignright"><strong>Importe de la<br />transacción</strong></th>
        

        <th valign="middle" class="alignright"><strong>Cantidad neta&nbsp;</strong></th>

        
      </tr>
    </thead>
    <tbody>

      
      
      <tr>
        <td valign="top"><span>26 octubre 2015</span></td>
        <td valign="top" style="90px"><span>505176946672480256</span></td>
        <td valign="top"><span>

  
    Cuenta Hailo
  



</span></td>
        <td valign="top" style="width:20px"><span>1:47</span></td>

        <td valign="top"><span>Carrer dels Almogàvers El Parc i la Llacuna del Poblenou -<br />Ronda de Sant Antoni 59 08011 Barcelona</span></td>

        
        <td valign="top" class="alignright"><span>7,00</span></td>
        

        <td valign="top" class="alignright"><span>-0,85</span></td>

        <td valign="top" class="alignright"><span>7,00</span></td>

        <td valign="top" class="alignright"><span>6,15</span></td>
      </tr>
      
      
      
      <tr>
        <td valign="top"><span>26 octubre 2015</span></td>
        <td valign="top" style="90px"><span>505186423706738688</span></td>
        <td valign="top"><span>

  
    Cuenta Hailo
  



</span></td>
        <td valign="top" style="width:20px"><span>2:34</span></td>

        <td valign="top"><span>Carrer dels Agudells El Carmelo -<br />Carrer de Ramon Llull 480 08930 Sant Adrià de Besòs</span></td>

        
        <td valign="top" class="alignright"><span>17,80</span></td>
        

        <td valign="top" class="alignright"><span>-2,15</span></td>

        <td valign="top" class="alignright"><span>17,79</span></td>

        <td valign="top" class="alignright"><span>15,65</span></td>
      </tr>
      
      
      
      <tr>
        <td valign="top"><span>27 octubre 2015</span></td>
        <td valign="top" style="90px"><span>505858773036666880</span></td>
        <td valign="top"><span>

  
    Efectivo Hailo
  



</span></td>
        <td valign="top" style="width:20px"><span>11:03</span></td>

        <td valign="top"><span>Carrer Llobregat L&#39;Hospitalet de Llobregat -<br />Passeig de la Zona Franca 217 08038 Barcelona</span></td>

        
        <td valign="top" class="alignright"><span>6,30</span></td>
        

        <td valign="top" class="alignright"><span>-0,76</span></td>

        <td valign="top" class="alignright"><span>0,00</span></td>

        <td valign="top" class="alignright"><span>-0,76</span></td>
      </tr>
      
      
      
      <tr>
        <td valign="top"><span>28 octubre 2015</span></td>
        <td valign="top" style="90px"><span>505877274501681152</span></td>
        <td valign="top"><span>

  
    Cuenta Hailo
  



</span></td>
        <td valign="top" style="width:20px"><span>12:18</span></td>

        <td valign="top"><span>Passatge de Bosch i Labrús El Poblenou -<br />Carrer de Còrsega 255 08036 Barcelona</span></td>

        
        <td valign="top" class="alignright"><span>12,71</span></td>
        

        <td valign="top" class="alignright"><span>-1,54</span></td>

        <td valign="top" class="alignright"><span>12,71</span></td>

        <td valign="top" class="alignright"><span>11,17</span></td>
      </tr>
      
      
      
      <tr>
        <td valign="top"><span>28 octubre 2015</span></td>
        <td valign="top" style="90px"><span>505885074674409472</span></td>
        <td valign="top"><span>

  
    Cuenta Hailo
  



</span></td>
        <td valign="top" style="width:20px"><span>12:39</span></td>

        <td valign="top"><span>Carrer de Roger de Flor El Fort Pienc -<br />Carrer de Lope de Vega 259 08018 Barcelona</span></td>

        
        <td valign="top" class="alignright"><span>8,00</span></td>
        

        <td valign="top" class="alignright"><span>-0,97</span></td>

        <td valign="top" class="alignright"><span>8,00</span></td>

        <td valign="top" class="alignright"><span>7,03</span></td>
      </tr>
      
      
      
      <tr>
        <td valign="top"><span>30 octubre 2015</span></td>
        <td valign="top" style="90px"><span>506645991473274880</span></td>
        <td valign="top"><span>

  
    Efectivo Hailo
  



</span></td>
        <td valign="top" style="width:20px"><span>3:06</span></td>

        <td valign="top"><span>Carrer de Carreras i Candi Sants-Badal -<br />Travessera de Gràcia 41 08021 Barcelona</span></td>

        
        <td valign="top" class="alignright"><span>7,35</span></td>
        

        <td valign="top" class="alignright"><span>-0,89</span></td>

        <td valign="top" class="alignright"><span>0,00</span></td>

        <td valign="top" class="alignright"><span>-0,89</span></td>
      </tr>
      
      
      
      <tr>
        <td valign="top"><span>30 octubre 2015</span></td>
        <td valign="top" style="90px"><span>506673562659315712</span></td>
        <td valign="top"><span>

  
    Efectivo Hailo
  



</span></td>
        <td valign="top" style="width:20px"><span>5:00</span></td>

        <td valign="top"><span>Carrer de Viladomat La Nova Esquerra de l&#39;Eixample -<br />C-32B 08820 El Prat de Llobregat</span></td>

        
        <td valign="top" class="alignright"><span>24,35</span></td>
        

        <td valign="top" class="alignright"><span>-2,95</span></td>

        <td valign="top" class="alignright"><span>0,00</span></td>

        <td valign="top" class="alignright"><span>-2,95</span></td>
      </tr>
      
      

    </tbody>
  </table>
</div><!-- end jobs -->

<div class="clear"></div>





<!--

SUMMARY

 -->



<div class="hr">
  <h2>Exclusivamente para fines administrativos</h2>
</div>

<table class="align-top">
    <tr>
        <td>


<div class="payment-details">
    <table>
      <caption>Detalles del pago</caption>
      <tr>
        <td><span><strong class="heading">Número de cuenta:</strong></span></td>
        <td>
          <span>
            <strong class="heading">Transferencia realizada el:</strong>
            
            
          </span>
        </td>
      </tr>
      <tr>
        <td><span style="padding-right:10px">****-*****-***-****1088</span></td>
        <td><span style="padding-right:10px">5 noviembre 2015</span></td>
      </tr>
      <tr>
        <td class="note"><span class="deemphasized">Los datos bancarios se pueden cambiar en la app de conductores Hailo.</span></td>
        <td class="note"><span class="deemphasized">Dependiendo de tu banco, el ingreso puede tardar unos días en llegar.</span></td>
      </tr>
    </table>
</div> <!-- end payment-details -->


</td>
<td class="td-empty-space"></td>
<td>


<div class="summary">
  <table>
    <thead>
          <tr>
            <th valign="middle"></th>
            <th style="font-size:12px" valign="middle"><strong>Importe<br />(IVA incl.)</strong></th>
            <th style="font-size:12px" valign="middle"><strong>Coste Hailo<br />(IVA incl.)</strong></th>
            <th style="font-size:12px" valign="middle"><strong>Neto</strong></th>
          </tr>
    </thead>
    
  <tbody>
  
    
  
  <tr>
    <td valign="top" class="first" style="font-size:12px"><span><strong>Hailo Efectivo</strong></span></td>
    <td valign="top" class="alignright"><span>38,00 €</span></td>
      <!-- gross - net round half down -->              
    <td valign="top" class="alignright"><span>-4,60 €</span></td>
    <td valign="top" class="alignright"><span>33,40 €</span></td>
  </tr>
  
  
  
  
  
  <tr>
    <td valign="top" class="first" style="font-size:12px"><span><strong>Hailo Cuenta y Tarjeta</strong></span></td>
    <td valign="top" class="alignright"><span>45,50 €</span></td>
      <!-- gross - net round half down -->              
    <td valign="top" class="alignright"><span>-5,51 €</span></td>
    <td valign="top" class="alignright"><span>39,99 €</span></td>
  </tr>
  
  

  
  
  <tr>
    <td valign="top" class="first" style="font-size:12px"><span><strong>Carreras no Hailo Tarjeta</strong></span></td>
    <td valign="top" class="alignright"><span>0,00 €</span></td>
      <!-- gross - net round half down -->              
    <td valign="top" class="alignright"><span>0,00 €</span></td>
    <td valign="top" class="alignright"><span>0,00 €</span></td>
  </tr>
  
  
  
  
  
  <tr>
    <td valign="top" class="first" style="font-size:12px"><span><strong>Subtotal</strong></span></td>
    <td valign="top" class="alignright"><span></span></td>
      <!-- gross - net round half down -->              
    <td valign="top" class="alignright"><span></span></td>
    
        
    <td valign="top" class="alignright"><span>73,39 €</span></td>
    
  </tr>
  
  <tr>
    <td>&nbsp;</td>
    <td></td>       
    <td></td>
    <td></td>
  </tr>  

  <tr>
    <td valign="top" class="first"><span><strong>Efectivo ya recibido</strong></span></td>
    <td valign="top" class="alignright"><span></span></td>
      <!-- gross - net round half down -->              
    <td valign="top" class="alignright"><span></span></td>
    <td valign="top" class="alignright"><span>
    
    -
    
    38,00 €
    </span></td>
  </tr>
  
  <tr>
    <td valign="top" class="first"><span><strong>Pago Hailo al conductor</strong></span></td>
    <td valign="top" class="alignright"><span></span></td>
      <!-- gross - net round half down -->              
    <td valign="top" class="alignright"><span></span></td>
    <td valign="top" class="alignright"><span><strong>35,40 €</strong></span></td>
  </tr>
  </tbody>
  </table>
</div><!-- end summary -->
<div class="clear"></div>

</td>
</tr>
</table>





<div class="extra">
    <h3>Echa un vistazo al Blog de Hailo en <a href="https://www.hailocab.com/barcelona/drivers/blog/">hailocab.com/barcelona/drivers/blog/</a>
      <br /><br />
    Necesitas ayuda? Contáctanos en <a href="mailto:barcelona.conductores@hailocab.com">barcelona.conductores@hailocab.com</a> [BCN6061]</h3>
</div>











  <div class="hr"></div>

  <div class="summary" style="padding:10px">
    <table>
      
    </table>
  </div>



<div class="clear"></div>




<p class="pagebreakhere"></p>

<div class="hr">
      <h2>Datos Factura</h2>
    </div>
    
    <div class="factura">
      <table>

          
        <caption>Concepto: Servicios prestados entre 26-10-15 y 01-11-15</caption>
        <tbody>
        
                    <tr>
            <td valign="top" rowspan=10><span><strong>FACTURA N&ordm;: Barcelona-6061-WJOKE_ww15<br/>&nbsp;<br/>FECHA FACTURA: Eeee, 01-11-15</strong><br/>&nbsp;<br/>Hailo Network Iberia SLU<br/>
B-86625951<br/>
C/ Santa Engracia 108, 1&ordm;D <br/>
28003 MADRID</span></td>
            <td valign="top"><span><strong>Conductor</strong></span></td>   
            <td valign="top"
            ><span>Evgeny Syrtsov</span></td>
                    </tr>
                    <tr>
            <td valign="top" class="first"><span><strong>DNI/NIE/NIF</strong></span></td>  
        <td valign="top"><span>x5307831Y</span></td>
                    </tr>
                    <tr>

              <td valign="top" class="first"><span><strong>No Licencia / No Carnet de Conductor</strong></span></td> <td valign="top"><span>53775</span></td>
                    </tr>
                    <tr>
                    
            <td valign="top" class="first"><span><strong>Direcci&oacute;n</strong></span></td>
                          <td valign="top">
                              <span>
                              
                                  Av . Paralelo 151, 1-3 <br />
                              
                              
                              08004 Barcelona 
                            </span>
                          </td>
                    </tr>
                    <tr>
                    
            <td valign="top" class="first"><span><strong>&nbsp;</strong></span></td>
                          <td valign="top" class="alignright">&nbsp;</td>
                    </tr>
                    
             <tr>
              <td valign="top" class="first"><span>Base imponible</span></td>
              <td valign="top"><span>-8,36 €</span></td>
            </tr>
             <tr>
              <td valign="top" class="first"><span>IVA</span></td>
              <td valign="top"><span>-1,75 €</span></td>
            </tr>
                      <tr>
              <td valign="top"><span><strong>Total Factura</strong></span></td>
              <td valign="top"><strong><span>10,11 €</span></strong></td>
            </tr>
            
                  </tbody>
      </table>
    </div><!-- end jobs -->

  </div><!-- end wrapper -->





</div><!-- end wrapper -->
  
 <!-- END SUMMARY -->


</body>
</html>


`

func TestBigTest(t *testing.T) {
	t.Logf("Num. proc.s: %d", runtime.GOMAXPROCS(-1))
	pdf, err := ConvertHtmlStringToPdf(htmlPage)
	if err != nil {
		t.Errorf("Error converting HTML to PDF: %s", err)
	}
	if len(pdf) != 38454 {
		t.Errorf("Wrong size for PDF output: expected: %d got: %d", 38454, len(pdf))
	}

	fName := fmt.Sprintf("%d.pdf", time.Now().Unix())
	wd, _ := os.Getwd()
	t.Logf("Open file for writing (wd: %s)... %s", wd, fName)
	f, err := os.OpenFile(fName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		t.Errorf("Failed to open file: %s\n", err)
	}
	defer func() { f.Close(); t.Logf("Closed PDF file: %s", fName) }()
	f.Truncate(0)
	f.Write([]byte(pdf))
}

// Is the problem something to do with running in a goroutine
func TestBigTestAsync(t *testing.T) {
	c := make(chan bool, 1)
	go func() {
		TestBigTest(t)
		c <- true
	}()
	<-c
}
