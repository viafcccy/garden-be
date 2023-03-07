po

持久对象 (persistent object)，po(persistent object) 就是在 Object/Relation Mapping 框架中的 Entity，po 的每个属性基本上都对应数据库表里面的某个字段。完全是一个符合 Java Bean 规范的纯 Java 对象，没有增加别的属性和方法。持久对象是由 insert 数据库创建，由数据库 delete 删除的。基本上持久对象生命周期和数据库密切相关。