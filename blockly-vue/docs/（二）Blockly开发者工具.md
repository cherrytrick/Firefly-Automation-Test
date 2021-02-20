# （二）Blockly开发者工具

[Blockly Developer Tools![img](https://developers.google.cn/blockly/images/external.svg)](https://blockly-demo.appspot.com/static/demos/blockfactory/index.html) 是基于Web的开发人员工具，可自动执行Blockly配置过程的各个部分，包括创建自定义块，构建工具箱和配置Web Blockly工作区。

使用Blockly developer 工具开发的三个流程：

- 使用Block Factory和Block Exporter创建块。
- 使用Workspace Factory创建一个工具箱和默认工作区。
- 使用Workspace Factory (当前只有web)配置你的工作区。

## Block Factory Tab（块工厂选项卡）

Block Factory tab帮助你为了自定义块创建 [block definitions](https://developers.google.cn/blockly/guides/create-custom-blocks/define-blocks)和 [code generators](https://developers.google.cn/blockly/guides/create-custom-blocks/generating-code). 在这个选项卡上，你可以轻松的创建、修改和保存自定义块。

### Defining a block（定义一个块）

该视频详细介绍了定义块的步骤。这个UI是过时的，但是这个块的功能描述仍然正确。



**中国大陆地区不可用**

https://developers.google.cn/_static/images/video-placeholder.svg



### Managing the library（管理依赖）

Blocks是通过名称来引用的，所以块必须有一个唯一的命名. 当你创建一个新快或者修改一个现有的块，UI会强制执行。

![img](https://developers.google.cn/blockly/images/block_save_as.png) ![img](https://developers.google.cn/blockly/images/update_button.png)

通过单击Library 按钮，你可以在之前保存的块之间切换或者创建一个新的空块。 更改一个已经存在的块的名称是一种其他快速创建类似定义的块的方法。

![img](https://developers.google.cn/blockly/images/blocklib_button.png)

### Exporting and importing a library（导出和导入一个依赖）

Blocks是保存到浏览器本地缓存中的.清理浏览器缓存意味着你将删除块. 想要永久保存你的快，你必须下载它的依赖. 你的Block Library 是一个已经下载的XML文件，你可以通过导入Block Library去设置你当前依赖的状态. 导入一个Block Library会替换掉现有的Block Library，所以首先应该先导出当前的Block Library.

推荐使用导入和导出功能来维护和共享不同的自定义块集。

![img](https://developers.google.cn/blockly/images/block_manage_buttons.png)

## Block Exporter tab（块导出选项卡）

设计完块之后，在应用程序中使用它们将需要导出块定义和生成存根。这是在“Block Exporter”选项卡上完成的。

任何一个块都存储在你的Block Library并且可以在Block Selector展示. 在这个块上单击去选择或者取消它的导入。如果你想要在依赖里选择全部的块，使用 “Select” → “All Stored In Block Library” 选项. 如果你创建了工具箱或者配置了你的工作区使用 Workspace Factory选项卡，你可以同样的使用点击“Select” → “All Used In Workspace Factory”来选择所有的块.

![img](https://developers.google.cn/blockly/images/block_exporter_select.png)

导出设置使您可以选择要定位的生成语言，以及是否要为所选块定义，生成器存根或两者。选择这些文件后，单击“Export”以下载文件。

![img](https://developers.google.cn/blockly/images/block_exporter_tab.png)

**注意:** 如果你使用Mac，保存对话框只能在同一时间保存一个文件 [one file at a time](https://github.com/google/blockly/issues/647)

## Workspace Factory tab（工作区工厂选项卡）

通过工作区工厂，可以轻松配置工作区中的工具箱和默认块集。您可以使用“Toolbox”和“Workspace”按钮在编辑工具箱和起始工作区之间切换。

![img](https://developers.google.cn/blockly/images/ws_fac_tb_ws_buttons.png)

### Building a toolbox（创建一个工具箱）

此选项卡有助于构建工具箱的XML。该材料假定您熟悉 [Toolbox](https://developers.google.cn/blockly/guides/configure/web/toolbox). 果已经具有要在此处编辑的工具箱的XML，则可以通过单击"Load to Edit"来加载它。



### Toolbox without categories（没有类别的工具箱）

如果您有几个块并且要显示它们而没有任何类别，只需将它们拖到工作区中，您将看到您的块出现在预览的工具箱中。

![img](https://developers.google.cn/blockly/images/workspace_fac_no_cat.png)

### Toolbox with categories（有类别的工具箱）

如果要按类别显示块，请单击“ +”按钮，然后为新类别选择下拉项。这会将类别添加到您可以选择和编辑的类别列表中。选择“标准类别”以添加单个标准的块类别（逻辑，循环等），或选择“标准工具箱”以添加所有标准的块类别。使用箭头按钮重新排序类别。

![img](https://developers.google.cn/blockly/images/category_menu.png)

**Note:** The standard categories and toolbox include all the blocks in the [Playground](https://blockly-demo.appspot.com/static/tests/playground.html). This set of blocks is not appropriate for most apps and should be pruned as needed. Also, some blocks are not supported on mobile yet.

**注意：**标准类别和工具箱包括[Playground](https://blockly-demo.appspot.com/static/tests/playground.html)中的所有块 。这组块不适用于大多数应用程序，应根据需要进行修剪。此外，移动设备尚不支持某些功能块。

To change the selected category’s name or color use the “Edit Category” dropdown. Dragging a block into the workspace will add it to the selected category.

要更改所选类别的名称或颜色，请使用“编辑类别”下拉菜单。将块拖到工作区中会将其添加到所选类别。

![img](https://developers.google.cn/blockly/images/edit_category.png)

### Advanced blocks（高级块）

By default, you can add any of the standard blocks or any blocks in your library to the toolbox. If you have blocks defined in JSON that aren't in your library, you can import them using the "Import Custom Blocks" button.

Some blocks should be used together or include defaults. This is done with [groups and shadows](https://developers.google.cn/blockly/guides/configure/web/toolbox#block_groups). Any blocks that are connected in the editor will be added to the toolbox as a group. Blocks that are attached to another block can also be changed to shadow blocks by selecting the child block and clicking the "Make Shadow" button. Note: Only child blocks that don't contain a variable may be changed to shadow blocks.

If you include a variable or function block in their toolbox, include a “Variables” or “Functions” category in your toolbox to allow users to fully utilize the block. Learn more about [“Variables” or “Functions" categories](https://developers.google.cn/blockly/guides/configure/web/toolbox#categories).

默认情况下，您可以将任何标准块或库中的任何块添加到工具箱。如果您的库中没有使用JSON定义的块，则可以使用“导入自定义块”按钮导入它们。

某些块应一起使用或包括默认值。这是通过 [组和阴影完成的](https://developers.google.cn/blockly/guides/configure/web/toolbox#block_groups)。在编辑器中连接的所有块都将作为一个组添加到工具箱中。通过选择子块并单击“制作阴影”按钮，还可以将连接到另一个块的块更改为阴影块。注意：只有不包含变量的子块才可以更改为阴影块。

如果在其工具箱中包含变量或功能块，请在工具箱中包含“变量”或“功能”类别，以允许用户充分利用该块。了解有关 [“变量”或“功能”类别的更多信息](https://developers.google.cn/blockly/guides/configure/web/toolbox#categories)。

### Configuring a workspace (for web Blockly)配置一个工作区

To configure different parts of your workspace, go to the “Workspace Factory” tab and select “Workspace".

要配置工作区的不同部分，请转到“工作区工厂”选项卡，然后选择“工作区”。

#### Choose Workspace Options

Set different values for [configuration options](https://developers.google.cn/blockly/guides/get-started/web#configuration) and see the result in the preview area. Enabling [grid](https://developers.google.cn/blockly/guides/configure/web/grid) or [zoom](https://developers.google.cn/blockly/guides/configure/web/zoom) reveals more options to configure. Also, switching to using categories usually requires a more complex workspace; a trashcan and scrollbars are added automatically when you add your first category.

为[配置选项](https://developers.google.cn/blockly/guides/get-started/web#configuration)设置不同的值， 并在预览区域中查看结果。启用 [网格](https://developers.google.cn/blockly/guides/configure/web/grid)或 [缩放会](https://developers.google.cn/blockly/guides/configure/web/zoom)显示更多配置选项。同样，切换到使用类别通常需要更复杂的工作空间；添加第一个类别时，会自动添加垃圾桶和滚动条。

![img](https://developers.google.cn/blockly/images/configure_workspace.png)

#### Add Pre-loaded Blocks to the Workspace 添加预定义块到工作区

This is optional but may be necessary if you want to display a set of blocks in the workspace:

- When the application loads.
- When an event (advancing a level, clicking a help button, etc.) is triggered.

Drag blocks into the editing space to see them in your workspace in the preview. You can create block groups, disable blocks, and make certain blocks shadow blocks when you select them.

这是可选的，但如果要在工作空间中显示一组块，则可能是必需的：

- 应用程序加载时。
- 触发事件（前进级别，单击帮助按钮等）时。

将块拖到编辑空间中，以在预览中的工作区中查看它们。选择它们时，可以创建块组，禁用块，并使某些块成为阴影块。

![img](https://developers.google.cn/blockly/images/load_workspace_blocks.png)

You can export these blocks as XML (see below). Add them to your workspace with `Blockly.Xml.domToWorkspace`, immediately after you create your workspace:

您可以将这些块导出为XML（请参见下文）。`Blockly.Xml.domToWorkspace`创建工作区后，立即使用来将它们添加到您的工作 区中：



```
var xmlText = '<xml xmlns="https://developers.google.com/blockly/xml">' +    '<block type="math_number"></block></xml>';Blockly.Xml.domToWorkspace(Blockly.Xml.textToDom(xmlText), workspace);
```

This sample code adds a single `math_number` block to the workspace.

此示例代码`math_number`向工作空间添加了一个块。

### Exporting（导出）

Workspace Factory gives you the following export options:

Workspace Factory为您提供以下导出选项：

![img](https://developers.google.cn/blockly/images/workspace_export_opt.png)

- Starter Code: Produces starter html and javascript to inject your customized Blockly workspace.
- Toolbox: Produces XML to specify your toolbox.
- Workspace Blocks: Produces XML which can be loaded into a workspace.

- 入门代码：生成入门html和javascript以注入自定义的Blockly工作区。
- 工具箱：生成XML以指定您的工具箱。
- 工作区块：生成可以加载到工作区中的XML。