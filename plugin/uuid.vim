function! Generate_uuid() abort
  let l:uuid = system('uuid')
  execute "normal! a".uuid."\<esc>"
endfunction

nnoremap <silent><C-S-u> :call Generate_uuid()<CR>

command! -count -nargs=* UUID call Generate_uuid()
