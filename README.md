# 米粒之光自动化测试平台

## 1.引言

### 1.1 编写目的

#### 项目名称：米粒之光自动化测试平台（Firefly-Automation-Test【FAT】）

#### 项目成员

发起人：万能的苗

#### 设计初衷

在19年年初的时候，用python写了一个自动化测试。对selenium做了封装，提供了一个exe的UI界面。当时是通过读取配置文件，无需更改代码。后来，用在升级冒烟测试，使原本3个人3小时的工作量缩短到了1个人一台电脑几分钟。这是我第一次见识到自动化的威力。而我本人也是特别懒的那种，能自动从来不手动。也开发了很多自动化的工具，比如自动Git提交代码、jenkins编译前自动修改某些文件的变量等等。

20年底，准备打造一个无需测试懂代码，只需根据业务逻辑，在可视化拼图界面生成python自动化脚本的开源自动化测试平台。

帮助广大的开发、测试人员解放双手，拥有更多的时间去学习、娱乐是本项目开发的初衷。

### 1.2 项目背景

经过考察，没有一个使用go做后端的自动化测试平台开源项目。并且他们只是初步封装了selenium、appnium等框架，使用下拉框选项的形式生成脚本。无法满足更复杂的业务需求。

之前参加“链谷杯”区块链应用创新大赛，在决赛的时候初次接触到Google blockly可视化编程技术。这个拼图生成python代码的技术深深的吸引了我。感觉这个就是解决可视化编程问题的天生利器。

也正好赶上了学习go语言的节点。因为go天然支持并发和gprc通信，所以就打算用go做后端服务调用生成好的python脚本打造一个可视化编程的自动化测试平台。

此时，是实践go语言的最佳时机。

### 1.2 项目概要

本项目采取GitHub开源方式，使用Apache2.0开源协议。

GitHub地址：

## 2.总体要求

### 2.1 总体功能架构

![avatar](http://assets.processon.com/chart_image/5fc9ccc3e0b34d4f98d0e2c2.png)



### 2.2 总体技术架构

![avatar](http://assets.processon.com/chart_image/5fc894d17d9c082f448bbb5f.png)

### 2.3 系统功能概述

从测试用例管理到测试脚本执行，再到代码质量检测的一整套开源解决方案。

### 2.4 系统功能要求

- 支持测试用例、评审、报表等管理
- 支持分布式动态扩容
- blockly二次开发，支持生成python自动化测试脚本
- selenium等自动化框架二次封装
- 支持用户权限分级
- 支持结果以邮件形式发送、bug邮件推送开发人员
- 支持测试结果统计，生成图表
- 部署支持Linux和windows双环境（服务器、个人PC）
- ffagent自研分布式中间件
  - 主从选举
  - 消息队列
  - 异步通信
  - 心跳检测
  - 任务调度
  - 服务发现注册

## 3.软件开发

### 3.1 需求分析

- 需要满足测试无需会python，通过页面就能生成python脚本。
- 满足高性能、高可用的现代软件设计理念。
- 提供测试的全环节平台化管理
- 从开发到测试到部署，全流程代码质量监测

### 3.2 概要设计

#### 3.2.1 概要设计说明

- 每个节点都会有一个ffagent作为分布式协调组件。

- 脚本的分发和执行，都是通过ffagent实现Reactor模型来完成的。
- master节点上有所有类型的本地消息队列
- broker节点上有自己拥有角色类型的本地消息队列

#### 3.2.3 基本设计概念和处理流程

![avatar](http://assets.processon.com/chart_image/5fcdc9fcf346fb50b83409de.png)

![avatar](http://assets.processon.com/chart_image/5fcddb0ce401fd19980df1ec.png)

### 3.3 详细设计

- 优先度（高、中、低）
  - 低优先度有可能会在选型或发现问题后移除

#### 3.3.1 自动化脚本生成（高）

- 支持playwright浏览器录制脚本

- 二次开发blockly
- blockly参考文章
  - https://blog.csdn.net/benwdm/article/details/84910517?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-9.control&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-9.control
- playwright参考文章
  - https://github.com/microsoft/playwright-python
  - https://mp.weixin.qq.com/s/h-Jt1QDfYdx8Wmhmn8cuYg

#### 3.3.2 节点通信（低）

- go实现Reactor模型
- 使用grpc通信，protobuf通信协议

#### 3.3.3 服务注册发现（低）

- 使用Redis作为服务中介或使用go开发一个简易kv数据库作为服务中介
- go中文网有go实现一个Redis的例子

#### 3.3.4 本地消息队列（中）

- 每个节点中的ffagent内置一个本地消息队列，用来缓存和队列执行分发的任务

#### 3.3.5 任务调度（高）

- 由master主节点进行任务的分发调度
- 应充分考虑各节点资源的不同，使用加权算法分发任务
- 考虑节点资源不足时的拒绝策略

#### 3.3.6 心跳检测（低）

- 集成现有框架或参考keepalive模型自主实现

#### 3.3.7 压力测试（高）

- 使用go的多线程并发技术完成高性能的压力测试

#### 3.3.8 WEB测试（中）

- 充分封装python的selenium框架，以支持blockly生成脚本

#### 3.3.9 APP测试（高）

- 充分封装python的appnium框架，以支持blockly生成脚本

#### 3.3.10 接口测试（高）

- 支持多协议，多请求格式的自定义推送
- 参考swgger的支持形式
- 最好可以将swgger集成

#### 3.3.10 测试用例管理（高）

- 提供好用的管理页面
- 支持多权限分级管理和评审

#### 3.3.11 报表功能（中）

- 使用可视化图形界面
- 根据脚本返回的数据（会存入mysql）生成页面图表

#### 3.3.12 邮件发送（中）

- 自动化测试结果通知
- bug通知开发人员

#### 3.3.13 环境支持（中）

- 应当支持windows环境部署
- 应当支持云环境部署
- 应当支持linux和容器部署

#### 3.3.14 代码质量检测（低）

- sonorqube、findbugs、pmd的支持
- 可以根据阿里代码规约插件二次开发IDEA插件

#### 3.3.15 协议转换和请求报文校验（中）

- 建议支持跨协议请求比如dubbo转protobuf等
- 请求文本格式校验，比如json的校验，xml的校验等等

#### 3.3.16 WEB服务（中）

- 使用Gin 作为web服务
- 与blockly配合使用

#### 3.3.17 golang调用python脚本（高）

- golang使用命令行等方式调用python脚本
- 获取python脚本的执行结果
- 异步执行
- 参考文章
  - https://studygolang.com/articles/20011
  - https://blog.csdn.net/hjx_dou/article/details/108105895

#### 3.3.18 分发python脚本（高）

- 考虑如何在项目启动的时候，确保每个节点都有正确的脚本文件
- 防止重复分发
- 压缩传输

### 3.4 数据库设计

#### 3.4.1 数据库说明

计划支持mysql和postgresql双数据库。

#### 3.4.2 数据库操作

#### 3.4.3 数据库结构

#### 3.4.4 表结构