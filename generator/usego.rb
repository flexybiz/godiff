#!/usr/bin/env ruby

require 'ffi'

module GoUuid
    extend FFI::Library
      ffi_lib './gouuid/goUuid.so'
        attach_function :GoUuid, [], :string
end

puts GoUuid.GoUuid()

