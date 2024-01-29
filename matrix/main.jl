function main()
    bytes = 0
    while !eof(stdin)
        line = readline(stdin)
        if line=="q"
            break
        end
        bytes += length(line)
    end
    println(bytes)
end

main()