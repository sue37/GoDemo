module Test/ArchTest/main

go 1.15

replace (
	Test/ArchTest/ClassA => ../ClassA
	Test/ArchTest/ClassB => ../ClassB
	Test/ArchTest/Factory => ../Factory
	Test/ArchTest/Intf => ../Intf
)

require (
	Test/ArchTest/Factory v0.0.0-00010101000000-000000000000
	Test/ArchTest/Intf v0.0.0-00010101000000-000000000000
)
