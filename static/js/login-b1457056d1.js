!function(){function e(t,r,i){function a(o,c){if(!r[o]){if(!t[o]){var s="function"==typeof require&&require;if(!c&&s)return s(o,!0);if(n)return n(o,!0);var l=new Error("Cannot find module '"+o+"'");throw l.code="MODULE_NOT_FOUND",l}var u=r[o]={exports:{}};t[o][0].call(u.exports,function(e){return a(t[o][1][e]||e)},u,u.exports,e,t,r,i)}return r[o].exports}for(var n="function"==typeof require&&require,o=0;o<i.length;o++)a(i[o]);return a}return e}()({1:[function(require,module,exports){$(function(){function e(){window.collectTest("get_verification_click",{get_verification_source:"mobile_login"}),accountSDK.getMobileCode({mobile:c.val(),captcha:s.val(),type:24}).then(function(e){Auth.countDown(l,60),i.showSucTip("code"),ui.success("已发送")},function(t){1104===t.error_code?Auth.pointVerify(e):1105===t.error_code?Auth.slideVerify(e):(ui.error(t.message),t.captcha&&(o[n]=!0,Auth.setCaptcha(t.captcha)))})}function t(){accountSDK.userLogin({mobile:a.mobile.value,code:a.code.value,captcha:a.pcode.value}).then(function(e){window.collectTest("login_success",{channel:n}),Auth.doLogin(e.data)},function(e){1104===e.error_code?Auth.pointVerify(t):1105===e.error_code?Auth.slideVerify(t):(ui.error(e.message),e.captcha&&(o[n]=!0,Auth.setCaptcha(e.captcha)))})}function r(){accountSDK.userLogin({account:a.email.value,password:a.password.value,captcha:a.pcode.value},!0).then(function(e){window.collectTest("login_success",{channel:n}),Auth.doLogin(e.data)},function(e){1104===e.error_code?Auth.pointVerify(r):1105===e.error_code?Auth.slideVerify(r):(ui.error(e.message),e.captcha&&(o[n]=!0,Auth.setCaptcha(e.captcha)))})}window.collectTest("enter_login",{});var i=require("validatorReg"),a=document.getElementById("regform"),n="mobile",o={mobile:!1,password:!1};$(a).checkForm(),$(".js-verify-img").off().on("click",function(){Auth.refreshCaptcha()}),$(".js-change-login").off().on("click",function(){"mobile"==n?($(".js-account").removeClass("hide"),$(".js-mobile").addClass("hide"),$(this).text("手机验证码登录"),n="password"):"password"==n&&($(".js-mobile").removeClass("hide"),$(".js-account").addClass("hide"),$(this).text("密码登录（手机号或邮箱）"),n="mobile"),o[n]?$(".js-verify-img-wrap").removeClass("hide"):$(".js-verify-img-wrap").addClass("hide")});var c=$("#mobile"),s=$("#pcode"),l=$("#codeBtn");$("#code");ui.btn.disable(l),c.on("blur",function(){var e=$(this);/^\d{11}$/i.test(c.val())?(0==l.data("counting")&&ui.btn.enable(l),i.showSucTip(e.attr("name"))):(ui.btn.disable(l),i.showErr(e,0))}),l.on("click",function(t){t.preventDefault(),"disabled"!=$(this).attr("disabled")&&e()}),$("#regform input[type=submit]").off().on("click",function(e){e.preventDefault(),window.collectTest("login_click",{channel:n}),"mobile"==n?t():"password"==n&&r()}),$("#regform").on("click",".js-sina-login",function(e){e.preventDefault(),window.collectTest("login_click",{channel:"weibo"}),accountSDK.ssoLogin({platform_app_id:343,platform:"sina_weibo",next:location.origin+"/public/socialLogin"})}),$("#regform").on("click"," .js-qq-login",function(e){e.preventDefault(),window.collectTest("login_click",{channel:"qq"}),accountSDK.ssoLogin({platform_app_id:342,platform:"qzone_sns",next:location.origin+"/public/socialLogin"})}),$("#regform").on("click"," .js-wechat-login",function(e){e.preventDefault(),window.collectTest("login_click",{channel:"wechat"}),accountSDK.ssoLogin({platform_app_id:375,platform:"weixin",next:location.origin+"/public/socialLogin"})})})},{validatorReg:2}],2:[function(require,module,exports){var Common=new Object;Common.trim=function(e){return e.replace(/(^\s*)|(\s*$)/g,"")},Common.strlen=function(e){return"utf-8"==($.browser.msie?document.charset:document.characterSet).toLowerCase()?e.replace(/[\u4e00-\u9fa5]/g,"***").length:e.replace(/[^\x00-\xff]/g,"**").length},validator={errmsg:"errmsg",errorTip:"errorTip",require:/[^(^\s*)|(\s*$)]/,email:/^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/,mobile:/^\d{11}$/i,userName:/^[\u0391-\uFFE5A-Za-z0-9_-]*$/i,limit:"this.checkLimit(Common.strlen(value))",limitB:"this.checkLimit(this.LenB(value))",repeat:"this.checkRepeat(value)",ajax:"this.doajax(errindex)",nRepeat:"!this.checkRepeat(value)",checkLimit:function(e){var t=this.element.attr("min")||Number.MIN_VALUE,r=this.element.attr("max")||Number.MAX_VALUE;return t<=e&&e<=r},LenB:function(e){return e.replace(/[^\x00-\xff]/g,"**").length},checkRepeat:function(e){var t=this.element.attr("to"),r=$('input[name="'+t+'"]').eq(0).val();return r||(r=$("#"+t).val()),e==r},doajax:function(e){var t,r,i=(this.element,this.errindex),a=this.element.attr("url"),n=this.element.val(),o=this.element.attr("msg");o.indexOf("|")>-1?(t=o.split("|"),r=t[i]):r=o;var c=(this.element.attr("type"),this.element.attr("method")||"post"),s=this.element.attr("name");if(""==a||a==undefined)return ui.error("Please specify url"),!1;check_data=s+"="+n,validator.removeErr(this.element),this.element.parent("*").find("."+validator.errorTip).remove();var l=$.ajax({type:c,url:a,data:check_data,cache:!1,async:!1,success:function(e){return e=e.replace(/(^\s*)|(\s*$)/g,""),"success"!=e?(r=""==r?e:r,validator.showErrTip(s,r),!1):"success"==e?(validator.showSucTip(s),!0):void 0}}).responseText;return"success"==(l=l.replace(/(^\s*)|(\s*$)/g,""))}},validator.showErr=function(e,t){var r=e.attr("msg")||"unkonwn",i=r.split("|"),a=i[t]?i[t]:i[0],n=e.attr("name");return validator.showErrTip(n,a),!1},validator.showSucTip=function(e){$(".js-tip-"+e).addClass("hide"),$(".input[name="+e+"]").parents(".inputLine").addClass("success"),$("input[name="+e+"]").parents(".inputLine").removeClass("error")},validator.showErrTip=function(e,t){$("input[name="+e+"]").parents(".inputLine").addClass("error"),$(".input[name="+e+"]").parents(".inputLine").removeClass("success"),$(".js-tip-"+e).html(t).removeClass("hide")},validator.removeErr=function(e){e.parent("*").find("."+this.errorTip).remove()},validator.checkajax=function(element,datatype,errindex){var value=$.trim(element.val());return this.element=element,this.errindex=errindex,validator.removeErr(element),eval(this[datatype])},validator.checkDatatype=function(element,datatype){var value=$.trim(element.val());switch(this.element=element,validator.removeErr(element),datatype){case"repeat":case"limit":case"limitB":case"nRepeat":return eval(this[datatype]);default:return this[datatype].test(value)}},validator.check=function(e){var t=e.attr("datatype"),r=$.trim(e.val());if(void 0===t)return!0;if("true"!=e.attr("require")&&""==r)return!0;var i=t.split("|"),a=!0;return $.each(i,function(t,r){if("undefined"==typeof validator[r])return a=!1,!1;if("ajax"==r)return a=validator.checkajax(e,r,t);if(0==validator.checkDatatype(e,r))return validator.showErr(e,t),a=!1;var i=e.attr("name");validator.showSucTip(i)}),a},exports.showErr=validator.showErr,exports.showSucTip=validator.showSucTip,exports.showErrTip=validator.showErrTip,$.fn.extend({checkForm:function(){var e=$(this),t=e.find(":input[require]");t.blur(function(e){return validator.check($(this))}),e.submit(function(){var e=!0,r=new Array,i=0;return t.each(function(t){0==validator.check($(this))&&(e=!1,r[i++]=t)}),0!=e||(t.eq(r[0]).focus().select(),!1)})}})},{}]},{},[1]);